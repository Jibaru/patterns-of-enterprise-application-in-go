package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/jibaru/page-controller/internal/models"
)

type HomeController struct {
	Conn *sql.DB
}

func (c *HomeController) Handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Fetch posts from the database
	posts, err := models.GetAllPosts(c.Conn)
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
