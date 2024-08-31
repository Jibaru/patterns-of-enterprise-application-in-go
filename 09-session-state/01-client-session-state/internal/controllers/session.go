package controllers

import (
	"net/http"

	"github.com/jibaru/client-session-state/internal/models"
	"github.com/jibaru/client-session-state/internal/views"
)

type SessionController struct {
	view *views.SessionView
}

func NewSessionController() *SessionController {
	return &SessionController{
		view: views.NewSessionView(),
	}
}

func (c *SessionController) Home(w http.ResponseWriter, r *http.Request) {
	// Retrieve session state from cookies
	session := models.NewSessionFromRequest(r)

	// Render the home page
	c.view.RenderHome(w, session)
}

func (c *SessionController) Login(w http.ResponseWriter, r *http.Request) {
	// Simulate login process
	session := models.NewSession()
	session.UserID = "user123"
	session.Save(w)

	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusFound)
}

func (c *SessionController) Logout(w http.ResponseWriter, r *http.Request) {
	// Clear session
	models.ClearSession(w)

	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusFound)
}
