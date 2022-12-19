package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, req *http.Request) {
	// Check if there is a cookie
	c, err := req.Cookie("session_id")
	if err == nil {

		// If there is a cookie, check if there is a session for the
		// cookie id

	}
	fmt.Println(c)
}
