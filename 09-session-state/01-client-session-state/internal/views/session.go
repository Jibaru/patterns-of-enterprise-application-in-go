package views

import (
	"html/template"
	"net/http"

	"github.com/jibaru/client-session-state/internal/models"
)

type SessionView struct {
	templates *template.Template
}

func NewSessionView() *SessionView {
	templates := template.Must(template.ParseFiles("internal/views/templates/home.html"))
	return &SessionView{templates: templates}
}

func (v *SessionView) RenderHome(w http.ResponseWriter, session *models.Session) {
	data := map[string]interface{}{
		"UserID": session.UserID,
	}
	v.templates.ExecuteTemplate(w, "home.html", data)
}
