package handlers

import (
	"net/http"

	"github.com/samuelataklti/bookings/pkg/config"
	"github.com/samuelataklti/bookings/pkg/models"
	"github.com/samuelataklti/bookings/pkg/render"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	requestIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", requestIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//Perform some logic
	StringMap := make(map[string]string)
	StringMap["testing"] = "hello, again!"

	requestIp := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = requestIp

	//send data to template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: StringMap,
	})
}
