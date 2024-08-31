package controllers

import (
	"html/template"
	"net/http"

	"github.com/jibaru/server-session-state/internal/session"
)

type SessionHandler struct {
	sessionManager *session.SessionManager
	tmpl           *template.Template
}

func NewSessionHandler(sessionManager *session.SessionManager) *SessionHandler {
	tmpl := template.Must(template.ParseFiles("./internal/views/templates/home.html"))
	return &SessionHandler{
		sessionManager: sessionManager,
		tmpl:           tmpl,
	}
}

func (sh *SessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("session_id")
	if err != nil || sessionID.Value == "" {
		sessionID = &http.Cookie{
			Name:  "session_id",
			Value: sh.sessionManager.CreateSession(),
			Path:  "/",
		}
		http.SetCookie(w, sessionID)
	}

	switch r.URL.Path {
	case "/":
		sh.handleHome(w, r, sessionID.Value)
	case "/login":
		sh.handleLogin(w, r, sessionID.Value)
	case "/logout":
		sh.handleLogout(w, r, sessionID.Value)
	default:
		http.NotFound(w, r)
	}
}

func (sh *SessionHandler) handleHome(w http.ResponseWriter, _ *http.Request, sessionID string) {
	userID, _ := sh.sessionManager.GetSessionData(sessionID, "UserID")

	sh.tmpl.Execute(w, map[string]interface{}{
		"UserID": userID,
	})
}

func (sh *SessionHandler) handleLogin(w http.ResponseWriter, r *http.Request, sessionID string) {
	if r.Method == http.MethodPost {
		sh.sessionManager.SetSessionData(sessionID, "UserID", "User123")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (sh *SessionHandler) handleLogout(w http.ResponseWriter, r *http.Request, sessionID string) {
	sh.sessionManager.DeleteSession(sessionID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
