package common

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/reivaj05/GoJSON"
	requester "github.com/reivaj05/GoRequester"
)

const (
	LayoutTemplate = "common/templates/layout.html"
	ApiURL         = "http://localhost:8000/"
)

var requests = requester.New()

func MakeRequest(body, method, endpoint string) (string, int, error) {
	return requests.MakeRequest(createRequestConfig(body, method, endpoint))
}

func createRequestConfig(body, method, endpoint string) *requester.RequestConfig {
	return &requester.RequestConfig{
		Body:    []byte(body),
		Method:  method,
		Headers: map[string]string{"Content-Type": "application/json"},
		URL:     fmt.Sprintf("%s%s", ApiURL, endpoint),
	}
}

func SetAuthCookies(rw http.ResponseWriter, response string) {
	http.SetCookie(rw, &http.Cookie{
		Name:  "loggedIn",
		Value: "true",
		Path:  "/",
	})
	http.SetCookie(rw, &http.Cookie{
		Name:  "userID",
		Value: getUserID(response),
		Path:  "/",
	})
}

func ClearAuthCookies(rw http.ResponseWriter) {
	http.SetCookie(rw, &http.Cookie{
		Name:   "loggedIn",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.SetCookie(rw, &http.Cookie{
		Name:   "userID",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}

func getUserID(response string) string {
	data, _ := GoJSON.New(response)
	id, _ := data.GetFloatFromPath("ID")
	return strconv.Itoa(int(id))
}
