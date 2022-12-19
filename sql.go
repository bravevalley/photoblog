package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type userlogin struct {
	Us string
	Ps string
}

func checkLogindata(username, password string) error {
	var activeuser userlogin

	// Prepare the query for continuous use
	stmt, err := DB.Prepare(`
	SELECT username, password
	FROM userlogin
	WHERE username = ?
	`)
	if err != nil  {
		return err
	}

	// Execute the query
	row := stmt.QueryRow(username)

	// Retrieve Data from Database
	err = row.Scan(&activeuser.Us, &activeuser.Ps); if err != nil {
		return err
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(activeuser.Ps), []byte(password))

	// Check if username and password is correct
	if username != activeuser.Us || err != nil {
		return fmt.Errorf("WrongInfo")
	}

	return nil
}