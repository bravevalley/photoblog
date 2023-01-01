package app

import (
	// "io"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"photo.blog/db"
	"photo.blog/session"

	"github.com/google/uuid"
)

type Handlers struct {
	Template *template.Template
	// Database  *sql.DB
	// Redis *redis.Client
}

func (h *Handlers) login(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "Got here")

	// Check if there is a cookie
	c, err := req.Cookie("session_id")
	if err == nil {

		// If there is a cookie, check if there is a session for the
		// cookie id
		if err = session.CheckLogin(c.Value); err == nil {
			// There is a cookie and the cookie is in sessions db
			http.Redirect(w, req, "/dashboard", http.StatusSeeOther)
			return
		}
	}

	// Check if it is a form submit
	if req.Method == http.MethodPost {

		us := req.FormValue("Username")
		pw := req.FormValue("Password")
		em := req.FormValue("Email")
		si := req.FormValue("SignUP")

		fmt.Println(si)

		if si == "Sign Up" {

			err = db.CreateUser(us, pw, em)
			if err != nil {
				if err.Error() == "UserExist" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				w.WriteHeader(http.StatusUnauthorized)

			}

			http.Redirect(w, req, "/login", http.StatusSeeOther)
			return
		}

		// Check if there is a user with the username and Password
		err = db.CheckLogindata(us, pw)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println(err)
			return
		}

		// Create a new session ID
		session_id := uuid.Must(uuid.NewRandom())
		c = &http.Cookie{
			Name:  "session_id",
			Value: session_id.String(),
		}

		// Set the session ID in the Session Database
		if err = session.CreateSession(c.Value, us); err != nil {
			log.Fatalln("Cant create new user session; ", err)
		}

		// Send the new session Id to the client
		http.SetCookie(w, c)

		// Redirect to dashboard
		http.Redirect(w, req, "/dashboard", http.StatusSeeOther)
	}

	h.Template.ExecuteTemplate(w, "login.gohtml", nil)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
