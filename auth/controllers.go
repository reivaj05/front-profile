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
var resetPasswordTemplate = template.Must(template.ParseFiles(
	common.LayoutTemplate, "auth/templates/reset-password.html"))

type data struct {
	IsLogged bool
	HasError bool
	Success  bool
	Error    string
}

func signupHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		signup(rw, req)
	} else {
		signupTemplate.ExecuteTemplate(rw, "layout", data{IsLogged: false})
	}
}

func signup(rw http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	response, status, err := common.MakeRequest(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password),
		"POST", usersEndpoint)
	if status == http.StatusCreated && err == nil {
		common.SetAuthCookies(rw, response)
		http.Redirect(rw, req, "/profile/", http.StatusSeeOther)
	} else {
		signupTemplate.ExecuteTemplate(rw, "layout", data{IsLogged: false, HasError: true, Error: response})
	}
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		login(rw, req)
	} else {
		loginTemplate.ExecuteTemplate(rw, "layout", data{IsLogged: false})
	}
}

func login(rw http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	response, status, err := common.MakeRequest(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password), "POST", loginEndpoint)
	if status == http.StatusOK && err == nil {
		common.SetAuthCookies(rw, response)
		http.Redirect(rw, req, "/profile/", http.StatusSeeOther)
	} else {
		loginTemplate.ExecuteTemplate(rw, "layout", data{IsLogged: false, HasError: true, Error: response})
	}
}

func logoutHandler(rw http.ResponseWriter, req *http.Request) {
	common.ClearAuthCookies(rw)
	http.Redirect(rw, req, "/home/", http.StatusSeeOther)
}

func resetPasswordHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		resetPassword(rw, req)
		return
	}
	resetPasswordTemplate.ExecuteTemplate(rw, "layout", data{IsLogged: false})
}

func resetPassword(rw http.ResponseWriter, req *http.Request) {
	var d data
	email := req.FormValue("email")
	response, status, err := common.MakeRequest(fmt.Sprintf(`{"email": "%s"}`, email), "POST", resetPasswordEndpoint)
	if status == http.StatusOK && err == nil {
		d = data{IsLogged: false, Success: true}
	} else {
		d = data{IsLogged: false, HasError: true, Error: response}
	}
	resetPasswordTemplate.ExecuteTemplate(rw, "layout", d)
}
