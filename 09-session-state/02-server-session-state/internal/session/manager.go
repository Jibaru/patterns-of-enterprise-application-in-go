package session

import (
	"fmt"
	"sync"
	"time"
)

type SessionManager struct {
	sessions map[string]map[string]interface{}
	mu       sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]map[string]interface{}),
	}
}

func (sm *SessionManager) CreateSession() string {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sessionID := fmt.Sprintf("%v", time.Now().UnixMilli())
	sm.sessions[sessionID] = make(map[string]interface{})
	return sessionID
}

func (sm *SessionManager) GetSessionData(sessionID, key string) (interface{}, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	session, exists := sm.sessions[sessionID]
	if !exists {
		return nil, false
	}
	value, exists := session[key]
	return value, exists
}

func (sm *SessionManager) SetSessionData(sessionID, key string, value interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	session, exists := sm.sessions[sessionID]
	if !exists {
		session = make(map[string]interface{})
		sm.sessions[sessionID] = session
	}
	session[key] = value
}

func (sm *SessionManager) DeleteSession(sessionID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	delete(sm.sessions, sessionID)
}
