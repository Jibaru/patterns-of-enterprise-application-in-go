package main

import (
	"log"
	"net/http"

	"github.com/jibaru/model-view-controller/internal/controllers"
	"github.com/jibaru/model-view-controller/internal/db"
)

func main() {
	db, err := db.Setup()
	if err != nil {
		log.Fatalf("Failed to setup the database: %v", err)
	}
	defer db.Close()

	todoController := controllers.NewTodoController(db)

	http.HandleFunc("/todos", todoController.ListTodos)
	http.HandleFunc("GET /todos/create", todoController.ShowCreateTodoForm)
	http.HandleFunc("POST /todos/create", todoController.CreateTodo)
	http.HandleFunc("GET /todos/edit", todoController.ShowEditTodoForm)
	http.HandleFunc("POST /todos/edit", todoController.EditTodo)
	http.HandleFunc("/todos/delete", todoController.DeleteTodo)

	log.Println("Server started at http://localhost:8080/todos")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
