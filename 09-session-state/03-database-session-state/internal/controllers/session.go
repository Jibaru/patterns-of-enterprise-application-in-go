package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jibaru/database-session-state/internal/session"
)

const sessionCookieKey = "database_session_id"

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

func (sh *SessionHandler) createSessionCookie() (*http.Cookie, error) {
	sessID, err := sh.sessionManager.CreateSession()
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:  sessionCookieKey,
		Value: sessID,
		Path:  "/",
	}, nil
}

func (sh *SessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie(sessionCookieKey)
	if err != nil || sessionCookie.Value == "" {
		sessionCookie, err := sh.createSessionCookie()
		if err != nil {
			http.Error(w, "can not create session", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, sessionCookie)
	} else if sessionCookie.Value != "" {
		exists, err := sh.sessionManager.ExistsSession(sessionCookie.Value)
		if err != nil {
			http.Error(w, "can not verify session", http.StatusInternalServerError)
			return
		}

		if !exists {
			sessionCookie, err := sh.createSessionCookie()
			if err != nil {
				http.Error(w, "can not create session", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, sessionCookie)
		}
	}

	switch r.URL.Path {
	case "/":
		sh.handleHome(w, r, sessionCookie.Value)
	case "/login":
		sh.handleLogin(w, r, sessionCookie.Value)
	case "/logout":
		sh.handleLogout(w, r, sessionCookie.Value)
	default:
		http.NotFound(w, r)
	}
}

func (sh *SessionHandler) handleHome(w http.ResponseWriter, _ *http.Request, sessionID string) {
	data, err := sh.sessionManager.GetSessionData(sessionID)
	if err != nil {
		log.Println(err)
	}

	sh.tmpl.Execute(w, map[string]interface{}{
		"UserID": data.UserID,
	})
}

func (sh *SessionHandler) handleLogin(w http.ResponseWriter, r *http.Request, sessionID string) {
	if r.Method == http.MethodPost {
		sh.sessionManager.SetSessionData(sessionID, session.Data{
			UserID: "User123",
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (sh *SessionHandler) handleLogout(w http.ResponseWriter, r *http.Request, sessionID string) {
	sh.sessionManager.DeleteSession(sessionID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
