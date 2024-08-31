package session

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type SessionManager struct {
	db *sql.DB
}

type Data struct {
	UserID string `json:"user_id"`
}

func NewSessionManager(db *sql.DB) *SessionManager {
	return &SessionManager{db: db}
}

func (sm *SessionManager) CreateSession() (string, error) {
	sessionID := fmt.Sprintf("%v", time.Now().UnixMilli())

	query := "INSERT INTO sessions (session_id, data) VALUES (?, ?)"
	_, err := sm.db.Exec(query, sessionID, "{}")
	if err != nil {
		return "", err
	}

	return sessionID, nil
}

func (sm *SessionManager) GetSessionData(sessionID string) (*Data, error) {
	var data string
	query := "SELECT data FROM sessions WHERE session_id = ?"
	err := sm.db.QueryRow(query, sessionID).Scan(&data)
	if err != nil {
		return nil, err
	}

	var sessionData Data
	if err := json.Unmarshal([]byte(data), &sessionData); err != nil {
		return nil, err
	}

	return &sessionData, nil
}

func (sm *SessionManager) ExistsSession(sessionID string) (bool, error) {
	var data string
	query := "SELECT data FROM sessions WHERE session_id = ?"
	err := sm.db.QueryRow(query, sessionID).Scan(&data)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (sm *SessionManager) SetSessionData(sessionID string, sessionData Data) error {
	updatedData, err := json.Marshal(sessionData)
	if err != nil {
		return err
	}

	updateQuery := "UPDATE sessions SET data = ? WHERE session_id = ?"
	_, err = sm.db.Exec(updateQuery, string(updatedData), sessionID)
	return err
}

func (sm *SessionManager) DeleteSession(sessionID string) error {
	deleteQuery := "DELETE FROM sessions WHERE session_id = ?"
	_, err := sm.db.Exec(deleteQuery, sessionID)
	return err
}
