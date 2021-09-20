package handlers

import (
	"github.com/bworkman1/bookings/pkg/config"
	"github.com/bworkman1/bookings/pkg/models"
	"github.com/bworkman1/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) GeneralQuarters(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "general-quarters.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MajorsSuite(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "majors-suite.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ContactUs(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "contact-us.page.tmpl", &models.TemplateData{})
}
