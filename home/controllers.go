package home

import (
	"fmt"
	"html/template"
	"net/http"
)

var homeTemplate = template.Must(template.ParseFiles("home/home.html"))

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement home")

	rw.Header().Set("Content-Type", "text/html")
	homeTemplate.Execute(rw, nil)
}
