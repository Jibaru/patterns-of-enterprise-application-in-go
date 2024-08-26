package main

import (
	"log"
	"net/http"

	"github.com/jibaru/transform-view/internal/controllers"
)

func main() {
	http.HandleFunc("/", controllers.AlbumView)

	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
