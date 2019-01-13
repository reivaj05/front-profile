package home

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/reivaj05/front-profile/common"
)

var homeTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "home/templates/home.html"))

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement home")

	homeTemplate.ExecuteTemplate(rw, "layout", nil)
}
