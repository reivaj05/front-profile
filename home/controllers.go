package home

import (
	"html/template"
	"net/http"

	"github.com/reivaj05/front-profile/common"
)

var homeTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "home/templates/home.html"))

type data struct {
	IsLogged bool
}

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	homeTemplate.ExecuteTemplate(rw, "layout", data{IsLogged: isLogged(req)})
}

func isLogged(req *http.Request) bool {
	c, _ := req.Cookie("loggedIn")
	return c != nil
}
