package auth

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/reivaj05/front-profile/common"
)

var signupTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "auth/templates/signup.html"))
var loginTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "auth/templates/login.html"))
var logoutTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "auth/templates/logout.html"))
var resetPasswordTemplate = template.Must(template.ParseFiles(
	common.LayoutTemplate, "auth/templates/reset-password.html"))

func signupHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement signup")

	rw.Header().Set("Content-Type", "text/html")
	signupTemplate.ExecuteTemplate(rw, "layout", nil)
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement login")

	rw.Header().Set("Content-Type", "text/html")
	loginTemplate.ExecuteTemplate(rw, "layout", nil)
}

func logoutHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement logout")

	rw.Header().Set("Content-Type", "text/html")
	logoutTemplate.ExecuteTemplate(rw, "layout", nil)
}

func resetPasswordHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement reset pasword")

	rw.Header().Set("Content-Type", "text/html")
	resetPasswordTemplate.ExecuteTemplate(rw, "layout", nil)
}
