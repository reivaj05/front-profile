package profile

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/reivaj05/front-profile/common"
)

const (
	usersEndpoint = "/users/"
)

type profile struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Address   string `json:"Address"`
}

var profileTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "profile/templates/profile.html"))

func profileHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement profile")
	if req.Method == "POST" {
		updateProfile(req)
	}
	profileTemplate.ExecuteTemplate(rw, "layout", nil)
}

func updateProfile(req *http.Request) {
	fmt.Println("TODO: Implement profile PUT")
	req.ParseForm()
	response, status, err := common.MakeRequest(createProfileBody(req), "PUT", usersEndpoint+"1/")
	fmt.Println(response, status, err)
}

func createProfileBody(req *http.Request) string {
	data, _ := json.Marshal(&profile{
		Email:     req.FormValue("email"),
		FirstName: req.FormValue("firstName"),
		LastName:  req.FormValue("lastName"),
		Phone:     req.FormValue("phone"),
		Address:   req.FormValue("address"),
	})
	return string(data)
}
