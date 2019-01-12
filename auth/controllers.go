package auth

import (
	"fmt"
	"html/template"
	"net/http"
)

var homeTemplate = template.Must(template.ParseFiles("auth/templates/home.html"))
var signupTemplate = template.Must(template.ParseFiles("auth/templates/signup.html"))
var loginTemplate = template.Must(template.ParseFiles("auth/templates/login.html"))
var logoutTemplate = template.Must(template.ParseFiles("auth/templates/logout.html"))
var resetPasswordTemplate = template.Must(template.ParseFiles("auth/templates/reset-password.html"))

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement home")

	rw.Header().Set("Content-Type", "text/html")
	homeTemplate.Execute(rw, nil)
}

func signupHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement signup")

	rw.Header().Set("Content-Type", "text/html")
	signupTemplate.Execute(rw, nil)
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement login")

	rw.Header().Set("Content-Type", "text/html")
	loginTemplate.Execute(rw, nil)
}

func logoutHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement logout")

	rw.Header().Set("Content-Type", "text/html")
	logoutTemplate.Execute(rw, nil)
}

func resetPasswordHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement reset pasword")

	rw.Header().Set("Content-Type", "text/html")
	resetPasswordTemplate.Execute(rw, nil)
}
