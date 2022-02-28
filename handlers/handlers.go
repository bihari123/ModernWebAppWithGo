package handlers

import (
	"net/http"

	"example.com/app1/utilities"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utilities.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	utilities.RenderTemplate(w, "about.page.tmpl")
}
