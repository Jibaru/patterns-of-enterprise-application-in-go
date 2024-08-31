package main

import (
	"log"
	"net/http"

	"github.com/jibaru/remote-facade-server/internal/controllers"
)

func main() {
	http.HandleFunc("GET /address", controllers.GetAddressHandler)
	http.HandleFunc("POST /address", controllers.SetAddressHandler)

	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
