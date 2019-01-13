package common

import (
	"fmt"

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
