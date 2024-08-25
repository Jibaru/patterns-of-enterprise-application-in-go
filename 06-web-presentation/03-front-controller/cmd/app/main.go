package main

import (
	"log"
	"net/http"

	"github.com/jibaru/front-controller/internal/controllers"
	"github.com/jibaru/front-controller/internal/db"
)

func main() {
	// Initialize the database
	db, err := db.Setup()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Initialize the Front Controller
	frontController := controllers.NewFrontController()

	// Register commands
	frontController.RegisterCommand("/", &controllers.HomeCommand{Conn: db})
	frontController.RegisterCommand("/posts/create", &controllers.CreatePostCommand{Conn: db})
	frontController.RegisterCommand("/posts/edit", &controllers.EditPostCommand{Conn: db})

	// Start the server
	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", frontController); err != nil {
		log.Fatal("Server failed:", err)
	}
}
