package home

import (
	"github.com/reivaj05/front-profile/common"

	"github.com/reivaj05/GoServer"
)

var Endpoints = []*GoServer.Endpoint{
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/",
		Handler: common.WithHTMLContentType(homeHandler),
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/home/",
		Handler: common.WithHTMLContentType(homeHandler),
	},
}
