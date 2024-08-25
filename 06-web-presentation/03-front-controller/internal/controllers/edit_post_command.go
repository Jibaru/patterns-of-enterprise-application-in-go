package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/jibaru/front-controller/internal/models"
)

type EditPostCommand struct {
	Conn *sql.DB
}

func (epc *EditPostCommand) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		title := r.FormValue("title")
		content := r.FormValue("content")

		// Update the post in the database
		err := models.UpdatePost(epc.Conn, id, title, content)
		if err != nil {
			http.Error(w, "Failed to update post", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id := r.URL.Query().Get("id")
	post, err := models.GetPostById(epc.Conn, id)
	if err != nil {
		http.Error(w, "Failed to load post", http.StatusInternalServerError)
		return
	}

	// Render the edit post form
	tmpl, err := template.ParseFiles("./internal/views/templates/edit.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, post)
}
