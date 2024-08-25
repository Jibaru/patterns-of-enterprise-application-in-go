package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/jibaru/front-controller/internal/models"
)

type CreatePostCommand struct {
	Conn *sql.DB
}

func (cpc *CreatePostCommand) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")

		// Insert new post into the database
		err := models.CreatePost(cpc.Conn, title, content)
		if err != nil {
			http.Error(w, "Failed to create post", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Render the create post form
	tmpl, err := template.ParseFiles("./internal/views/templates/create.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
