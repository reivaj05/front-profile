package auth

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/reivaj05/front-profile/common"
)

const (
	resetPasswordEndpoint = "auth/reset_password/"
	loginEndpoint         = "auth/login/"
	usersEndpoint         = "users/"
)

var signupTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "auth/templates/signup.html"))
var loginTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "auth/templates/login.html"))
var logoutTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "auth/templates/logout.html"))
var resetPasswordTemplate = template.Must(template.ParseFiles(
	common.LayoutTemplate, "auth/templates/reset-password.html"))

func signupHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		signup(rw, req)
	} else {
		signupTemplate.ExecuteTemplate(rw, "layout", nil)
	}
}

func signup(rw http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	response, status, err := common.MakeRequest(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password), "POST", usersEndpoint)
	fmt.Println(response, status, err)
	// TODO: Handle response and errors
	http.Redirect(rw, req, "/profile/", http.StatusSeeOther)
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		login(rw, req)
	} else {
		loginTemplate.ExecuteTemplate(rw, "layout", nil)
	}
}

func login(rw http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	response, status, err := common.MakeRequest(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password), "POST", loginEndpoint)
	fmt.Println(response, status, err)
	// TODO: Handle response and errors
	http.Redirect(rw, req, "/profile/", http.StatusSeeOther)
}

func logoutHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement logout")

	logoutTemplate.ExecuteTemplate(rw, "layout", nil)
}

func resetPasswordHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement reset pasword")
	if req.Method == "POST" {
		resetPassword(req)
	}
	resetPasswordTemplate.ExecuteTemplate(rw, "layout", nil)
}

func resetPassword(req *http.Request) {
	fmt.Println("TODO: Implement reset pasword POST")
	email := req.FormValue("email")
	response, status, err := common.MakeRequest(fmt.Sprintf(`{"email": "%s"}`, email), "POST", resetPasswordEndpoint)
	fmt.Println(response, status, err)
}
