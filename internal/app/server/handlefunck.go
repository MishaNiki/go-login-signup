package server

import (
	"net/http"
	"time"
)

// handleLodin ...
func (serv *Server) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("session")
		logdedIn := (err != http.ErrNoCookie)
		if logdedIn {
			_, err := serv.session.FindUserID(c.Value)

			if err != nil {
				c.Expires = time.Now().AddDate(0, 0, -1)
				http.SetCookie(w, c)
				http.Redirect(w, r, "/login", http.StatusFound)
			}

			http.Redirect(w, r, "/home", http.StatusFound)
		}
	}
}

// handleSignup ...
func (serv *Server) handleSignup() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// handleForgod ...
func (serv *Server) handleForgod() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// handleHome ...
func (serv *Server) handleHome() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
