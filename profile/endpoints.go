package profile

import (
	"github.com/reivaj05/GoServer"
)

var Endpoints = []*GoServer.Endpoint{
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/profile/",
		Handler: profileHandler,
	},
}
