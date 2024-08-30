package main

import (
	"log"
	"net/http"

	"github.com/jibaru/application-controller/internal/controllers"
)

func main() {
	appController := &controllers.ApplicationController{}
	inputController := &controllers.InputController{AppController: appController}

	http.Handle("/action/", inputController)

	log.Println("Server started at http://localhost:8080")
	log.Println("Available routes")
	for _, action := range controllers.AvailableActions {
		log.Println(action)
	}
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
