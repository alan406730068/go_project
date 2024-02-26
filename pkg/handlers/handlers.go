package handlers

import (
	"net/http"

	"github.com/alan406730068/go_project/pkg/config"
	"github.com/alan406730068/go_project/pkg/models"
	"github.com/alan406730068/go_project/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.Appconfig
}

func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
