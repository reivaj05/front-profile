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
	signupEndpoint        = "auth/signup/"
)

var signupTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "auth/templates/signup.html"))
var loginTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "auth/templates/login.html"))
var logoutTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "auth/templates/logout.html"))
var resetPasswordTemplate = template.Must(template.ParseFiles(
	common.LayoutTemplate, "auth/templates/reset-password.html"))

func signupHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement signup")

	if req.Method == "POST" {
		signup(req)
	}

	signupTemplate.ExecuteTemplate(rw, "layout", nil)
}

func signup(req *http.Request) {
	fmt.Println("TODO: Implement signup POST")
	req.ParseForm()
	email := req.FormValue("email")
	password := req.FormValue("password")
	response, status, err := common.MakeRequest(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password), "POST", signupEndpoint)
	fmt.Println(response, status, err)
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement login")

	if req.Method == "POST" {
		login(req)
	}

	loginTemplate.ExecuteTemplate(rw, "layout", nil)
}

func login(req *http.Request) {
	fmt.Println("TODO: Implement login POST")
	req.ParseForm()
	email := req.FormValue("email")
	password := req.FormValue("password")
	response, status, err := common.MakeRequest(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password), "POST", loginEndpoint)
	fmt.Println(response, status, err)
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
	req.ParseForm()
	email := req.FormValue("email")
	response, status, err := common.MakeRequest(fmt.Sprintf(`{"email": "%s"}`, email), "POST", resetPasswordEndpoint)
	fmt.Println(response, status, err)
}
