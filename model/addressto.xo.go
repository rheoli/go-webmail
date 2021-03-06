// Package model contains the types for schema ''.
package model

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// AddressTo represents a row from 'address_to'.
type AddressTo struct {
	ID        int64 // id
	MailID    int64 // mail_id
	AddressID int64 // address_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AddressTo exists in the database.
func (at *AddressTo) Exists() bool {
	return at._exists
}

// Deleted provides information if the AddressTo has been deleted from the database.
func (at *AddressTo) Deleted() bool {
	return at._deleted
}

// Insert inserts the AddressTo to the database.
func (at *AddressTo) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if at._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO address_to (` +
		`mail_id, address_id` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, at.MailID, at.AddressID)
	res, err := db.Exec(sqlstr, at.MailID, at.AddressID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	at.ID = int64(id)
	at._exists = true

	return nil
}

// Update updates the AddressTo in the database.
func (at *AddressTo) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !at._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if at._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE address_to SET ` +
		`mail_id = ?, address_id = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, at.MailID, at.AddressID, at.ID)
	_, err = db.Exec(sqlstr, at.MailID, at.AddressID, at.ID)
	return err
}

// Save saves the AddressTo to the database.
func (at *AddressTo) Save(db XODB) error {
	if at.Exists() {
		return at.Update(db)
	}

	return at.Insert(db)
}

// Delete deletes the AddressTo from the database.
func (at *AddressTo) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !at._exists {
		return nil
	}

	// if deleted, bail
	if at._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM address_to WHERE id = ?`

	// run query
	XOLog(sqlstr, at.ID)
	_, err = db.Exec(sqlstr, at.ID)
	if err != nil {
		return err
	}

	// set deleted
	at._deleted = true

	return nil
}

// AddressByAddressID returns the Address associated with the AddressTo's AddressID (address_id).
//
// Generated from foreign key 'address_to_address_id_fkey'.
func (at *AddressTo) AddressByAddressID(db XODB) (*Address, error) {
	return AddressByID(db, at.AddressID)
}

// MailByMailID returns the Mail associated with the AddressTo's MailID (mail_id).
//
// Generated from foreign key 'address_to_mail_id_fkey'.
func (at *AddressTo) MailByMailID(db XODB) (*Mail, error) {
	return MailByID(db, at.MailID)
}

// AddressToByID retrieves a row from 'address_to' as a AddressTo.
//
// Generated from index 'address_to_id_pkey'.
func AddressToByID(db XODB, id int64) (*AddressTo, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, mail_id, address_id ` +
		`FROM address_to ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	at := AddressTo{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&at.ID, &at.MailID, &at.AddressID)
	if err != nil {
		return nil, err
	}

	return &at, nil
}

// AddressTosByAddressID retrieves a row from 'address_to' as a AddressTo.
//
// Generated from index 'idx_address_to_address_id'.
func AddressTosByAddressID(db XODB, addressID int64) ([]*AddressTo, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, mail_id, address_id ` +
		`FROM address_to ` +
		`WHERE address_id = ?`

	// run query
	XOLog(sqlstr, addressID)
	q, err := db.Query(sqlstr, addressID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AddressTo{}
	for q.Next() {
		at := AddressTo{
			_exists: true,
		}

		// scan
		err = q.Scan(&at.ID, &at.MailID, &at.AddressID)
		if err != nil {
			return nil, err
		}

		res = append(res, &at)
	}

	return res, nil
}

// AddressTosByMailID retrieves a row from 'address_to' as a AddressTo.
//
// Generated from index 'idx_address_to_mail_id'.
func AddressTosByMailID(db XODB, mailID int64) ([]*AddressTo, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, mail_id, address_id ` +
		`FROM address_to ` +
		`WHERE mail_id = ?`

	// run query
	XOLog(sqlstr, mailID)
	q, err := db.Query(sqlstr, mailID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AddressTo{}
	for q.Next() {
		at := AddressTo{
			_exists: true,
		}

		// scan
		err = q.Scan(&at.ID, &at.MailID, &at.AddressID)
		if err != nil {
			return nil, err
		}

		res = append(res, &at)
	}

	return res, nil
}
