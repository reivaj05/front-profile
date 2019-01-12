package home

import (
	"github.com/reivaj05/GoServer"
)

var Endpoints = []*GoServer.Endpoint{
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/",
		Handler: homeHandler,
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/home/",
		Handler: homeHandler,
	},
}
