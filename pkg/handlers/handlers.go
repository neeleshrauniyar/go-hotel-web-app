package handlers

import (
	"net/http"

	"github.com/neeleshrauniyar/go-hotel-web-app/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
	//fmt.Fprintf(w, "Hello Go")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}