package auth

import (
	"github.com/reivaj05/GoServer"
)

var Endpoints = []*GoServer.Endpoint{
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/signup/",
		Handler: signupHandler,
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/login/",
		Handler: loginHandler,
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/logout/",
		Handler: logoutHandler,
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/reset_password/",
		Handler: resetPasswordHandler,
	},
}
