package profile

import (
	"fmt"
	"html/template"
	"net/http"
)

var profileTemplate = template.Must(template.ParseFiles("profile/templates/profile.html"))

func profileHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement profile")

	rw.Header().Set("Content-Type", "text/html")
	profileTemplate.Execute(rw, nil)
}
