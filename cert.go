package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"path"
	"strings"
	"time"

	"github.com/nochso/mlog"
)

func prepareCert() {
	certPath := path.Join(dataDir, "cert.pem")
	keyPath := path.Join(dataDir, "key.pem")
	if _, err := os.Stat(certPath); err == nil {
		return
	}
	mlog.Info("Creating certificate")
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		mlog.Fatalf("failed to generate private key: %s", err)
	}
	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		mlog.Fatalf("failed to generate serial number: %s", err)
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"noch.so go-webmail"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	usedHosts := map[string]bool{}
	for _, h := range cfg.Accept.Domains {
		h = strings.Trim(h, ".")
		if usedHosts[h] {
			continue
		}
		usedHosts[h] = true
		if ip := net.ParseIP(h); ip != nil {
			mlog.Info("Adding IP address to certificate: %s", ip)
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			mlog.Info("Adding DNS name to certificate: %s", h)
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		mlog.Fatalf("Failed to create certificate: %s", err)
	}

	certOut, err := os.Create(certPath)
	if err != nil {
		mlog.Fatalf("failed to open '%s' for writing: %s", certPath, err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()
	mlog.Info("written certificate: %s", certPath)

	keyOut, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		mlog.Fatalf("failed to open %s for writing: %s", keyPath, err)
		return
	}
	pemBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)}
	pem.Encode(keyOut, pemBlock)
	keyOut.Close()
	mlog.Info("written private key: %s", keyPath)
}
