package main

import (
	"fmt"
	"strings"

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
	WHERE username=$1
	`)
	if err != nil {
		return err
	}
	
	// Execute the query
	row := stmt.QueryRow(username)

	// Retrieve Data from Database
	err = row.Scan(&activeuser.Us, &activeuser.Ps)
	if err != nil {
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

func createUser(Username, Password, Email string) error {
	ps, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Unable to encrypt password")
	}

	stmt, err := DB.Prepare(`
	INSERT INTO userlogin
	VALUES($1, $2, $3);
	`)

	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return fmt.Errorf("UserExist")
		}
		return fmt.Errorf("Can not prepare insert statement")
	}

	_, err = stmt.Exec(Username, string(ps), Email); if err != nil {
		return fmt.Errorf("Can not create user")
	}

	return nil
}