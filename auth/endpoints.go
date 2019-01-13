package auth

import (
	"github.com/reivaj05/front-profile/common"

	"github.com/reivaj05/GoServer"
)

var Endpoints = []*GoServer.Endpoint{
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/signup/",
		Handler: common.WithHTMLContentType(common.UserAuthenticated(signupHandler)),
	},
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/auth/signup/",
		Handler: common.WithHTMLContentType(common.UserAuthenticated(signupHandler)),
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/login/",
		Handler: common.WithHTMLContentType(common.UserAuthenticated(loginHandler)),
	},
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/auth/login/",
		Handler: common.WithHTMLContentType(common.UserAuthenticated(loginHandler)),
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/logout/",
		Handler: common.WithHTMLContentType(logoutHandler),
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/reset_password/",
		Handler: common.WithHTMLContentType(common.UserAuthenticated(resetPasswordHandler)),
	},
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/auth/reset_password/",
		Handler: common.WithHTMLContentType(common.UserAuthenticated(resetPasswordHandler)),
	},
}
