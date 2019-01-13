package common

import (
	"net/http"
)

func WithHTMLContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
		}
		w.Header().Set("Content-Type", "text/html")
		next.ServeHTTP(w, r)
	}
}

func UserAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := r.Cookie("loggedIn"); err == nil {
			http.Redirect(w, r, "/profile/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func UserUnauthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := r.Cookie("loggedIn"); err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/home/", http.StatusSeeOther)
				return
			}
		}
		next.ServeHTTP(w, r)
	}
}
