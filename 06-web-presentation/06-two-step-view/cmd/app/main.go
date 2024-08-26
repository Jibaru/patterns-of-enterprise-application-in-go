package main

import (
	"log"
	"net/http"

	"github.com/jibaru/two-step-view/internal/controllers"
)

func main() {
	// Route for viewing an album
	http.HandleFunc("/", controllers.AlbumView)

	// Start the server
	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
