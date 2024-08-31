package main

import (
	"log"
	"net/http"

	"github.com/jibaru/client-session-state/internal/controllers"
)

func main() {

	// Session controller
	sessionController := controllers.NewSessionController()

	// Routes
	http.HandleFunc("GET /", sessionController.Home)
	http.HandleFunc("POST /login", sessionController.Login)
	http.HandleFunc("GET /logout", sessionController.Logout)

	// Start the server
	log.Println("Server is running on port http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
