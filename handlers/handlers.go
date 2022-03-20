package handlers

import (
	"net/http"

	"example.com/app1/config"
	"example.com/app1/utilities"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	utilities.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	utilities.RenderTemplate(w, "about.page.tmpl")
}
