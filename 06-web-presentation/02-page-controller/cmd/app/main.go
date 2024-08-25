package main

import (
	"log"
	"net/http"

	"github.com/jibaru/page-controller/internal/controllers"
	"github.com/jibaru/page-controller/internal/db"
)

func main() {
	// Initialize the database
	db, err := db.Setup()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	homeCtrl := controllers.HomeController{Conn: db}
	postCtrl := controllers.PostController{Conn: db}

	// Route setup
	http.HandleFunc("/", homeCtrl.Handle)
	http.HandleFunc("/posts/create", postCtrl.CreatePost)
	http.HandleFunc("/posts/edit", postCtrl.EditPost)

	// Start the server
	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
