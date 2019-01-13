package profile

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/reivaj05/front-profile/common"
)

const (
	usersEndpoint = "users/"
)

type profile struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	ID        int    `json:"-"`
}

var profileTemplate = template.Must(template.ParseFiles(common.LayoutTemplate, "profile/templates/profile.html"))

type data struct {
	IsLogged bool
	User     profile
	Success  bool
}

func profileHandler(rw http.ResponseWriter, req *http.Request) {
	var success = false
	if req.Method == "POST" {
		updateProfile(req)
		success = true
	}
	profileTemplate.ExecuteTemplate(rw, "layout", data{IsLogged: true, User: getProfile(req), Success: success})
}

func updateProfile(req *http.Request) {
	id := getProfileID(req)
	endpoint := fmt.Sprintf("%s%d/", usersEndpoint, id)
	fmt.Println(common.MakeRequest(createProfileBody(req), "PUT", endpoint))
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

func getProfile(req *http.Request) profile {
	id := getProfileID(req)
	endpoint := fmt.Sprintf("%s%d/", usersEndpoint, id)
	response, _, _ := common.MakeRequest("", "GET", endpoint)
	return parseProfile(response)
}

func getProfileID(req *http.Request) int {
	c, _ := req.Cookie("userID")
	id, _ := strconv.Atoi(c.Value)
	return int(id)
}

func parseProfile(response string) profile {
	var pro profile
	json.Unmarshal([]byte(response), &pro)
	return pro
}
