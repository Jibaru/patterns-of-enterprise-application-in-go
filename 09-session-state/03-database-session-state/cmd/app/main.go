package main

import (
	"log"
	"net/http"

	"github.com/jibaru/database-session-state/internal/controllers"
	"github.com/jibaru/database-session-state/internal/db"
	"github.com/jibaru/database-session-state/internal/session"
)

func main() {
	dbConn, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sessionManager := session.NewSessionManager(dbConn)

	mux := http.NewServeMux()
	mux.Handle("/", controllers.NewSessionHandler(sessionManager))

	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
