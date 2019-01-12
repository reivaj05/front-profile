package profile

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/reivaj05/front-profile/common"
)

var profileTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "profile/templates/profile.html"))

func profileHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement profile")

	rw.Header().Set("Content-Type", "text/html")
	profileTemplate.ExecuteTemplate(rw, "layout", nil)
}
