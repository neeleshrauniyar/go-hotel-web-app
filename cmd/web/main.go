package main

import (
	"net/http"

	"github.com/neeleshrauniyar/go-hotel-web-app/pkg/handlers"
)

func main() {
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(":8080", nil)
}
