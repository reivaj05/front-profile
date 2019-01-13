package profile

import (
	"github.com/reivaj05/front-profile/common"

	"github.com/reivaj05/GoServer"
)

var Endpoints = []*GoServer.Endpoint{
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/profile/",
		Handler: common.IsUserLoggedIn(common.WithHTMLContentType(profileHandler)),
	},
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/profile/",
		Handler: common.IsUserLoggedIn(common.WithHTMLContentType(profileHandler)),
	},
}
