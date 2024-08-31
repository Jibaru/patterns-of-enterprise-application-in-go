package main

import (
	"log"
	"net/http"

	"github.com/jibaru/server-session-state/internal/controllers"
	"github.com/jibaru/server-session-state/internal/session"
)

func main() {
	sessionManager := session.NewSessionManager()

	http.Handle("/", controllers.NewSessionHandler(sessionManager))

	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
