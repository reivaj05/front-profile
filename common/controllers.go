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
