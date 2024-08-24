package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jibaru/model-view-controller/internal/models"
)

var templates = template.Must(template.ParseGlob("internal/views/templates/*.html"))

type TodoController struct {
	DB *sql.DB
}

func NewTodoController(db *sql.DB) *TodoController {
	return &TodoController{DB: db}
}

func (tc *TodoController) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAllTodos(tc.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = templates.ExecuteTemplate(w, "index.html", todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (tc *TodoController) ShowCreateTodoForm(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (tc *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		description := r.FormValue("description")

		err := models.CreateTodo(tc.DB, title, description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (tc *TodoController) ShowEditTodoForm(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := models.GetTodoByID(tc.DB, id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	err = templates.ExecuteTemplate(w, "edit.html", todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (tc *TodoController) EditTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		description := r.FormValue("description")
		completed := r.FormValue("completed") == "on"

		todo := models.Todo{
			ID:          id,
			Title:       title,
			Description: description,
			Completed:   completed,
		}

		err = models.UpdateTodo(tc.DB, todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (tc *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteTodo(tc.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}
