package auth

import (
	"github.com/reivaj05/front-profile/common"

	"github.com/reivaj05/GoServer"
)

var Endpoints = []*GoServer.Endpoint{
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/signup/",
		Handler: common.WithHTMLContentType(common.IsUserLoggedIn(signupHandler)),
	},
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/auth/signup/",
		Handler: common.WithHTMLContentType(common.IsUserLoggedIn(signupHandler)),
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/login/",
		Handler: common.WithHTMLContentType(common.IsUserLoggedIn(loginHandler)),
	},
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/auth/login/",
		Handler: common.WithHTMLContentType(common.IsUserLoggedIn(loginHandler)),
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/logout/",
		Handler: common.WithHTMLContentType(common.IsUserLoggedIn(logoutHandler)),
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/auth/reset_password/",
		Handler: common.WithHTMLContentType(common.IsUserLoggedIn(resetPasswordHandler)),
	},
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/auth/reset_password/",
		Handler: common.WithHTMLContentType(common.IsUserLoggedIn(resetPasswordHandler)),
	},
}
