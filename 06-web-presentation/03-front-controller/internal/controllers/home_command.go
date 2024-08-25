package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/jibaru/front-controller/internal/models"
)

type HomeCommand struct {
	Conn *sql.DB
}

func (hc *HomeCommand) Execute(w http.ResponseWriter, r *http.Request) {
	// Fetch posts from the database
	posts, err := models.GetAllPosts(hc.Conn)
	if err != nil {
		http.Error(w, "Failed to load posts", http.StatusInternalServerError)
		return
	}

	// Parse and execute the home template
	tmpl, err := template.ParseFiles("./internal/views/templates/home.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, posts)
}
