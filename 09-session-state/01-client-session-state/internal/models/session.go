package models

import (
	"net/http"
	"time"
)

// Session represents the session state
type Session struct {
	UserID string
}

// NewSession creates a new session instance
func NewSession() *Session {
	return &Session{}
}

// NewSessionFromRequest retrieves session data from cookies
func NewSessionFromRequest(r *http.Request) *Session {
	cookie, err := r.Cookie("session_user")
	if err != nil {
		return &Session{}
	}

	return &Session{
		UserID: cookie.Value,
	}
}

// Save stores the session data in a cookie
func (s *Session) Save(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session_user",
		Value:   s.UserID,
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	})
}

// ClearSession clears the session data by expiring the cookie
func ClearSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session_user",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})
}
