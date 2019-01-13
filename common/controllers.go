package common

import (
	"net/http"
	"strings"
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

func IsUserLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if c, err := r.Cookie("loggedIn"); err != nil {
			if err == http.ErrNoCookie && !strings.HasPrefix(r.RequestURI, "/auth/") {
				http.Redirect(w, r, "/home/", http.StatusSeeOther)
				return
			}
		} else if c != nil && strings.HasPrefix(r.RequestURI, "/auth/") {
			http.Redirect(w, r, "/profile/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}
