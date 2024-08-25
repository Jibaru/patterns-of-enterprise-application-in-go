package controllers

import (
	"html/template"
	"net/http"

	"github.com/jibaru/template-view/internal/domain"
	"github.com/jibaru/template-view/internal/views/helpers"
)

func RenderBook(w http.ResponseWriter, r *http.Request) {
	// Example book data
	book := domain.Book{
		Title:       "The Night Circus",
		Author:      "Erin Morgenstern",
		Description: "The Night Circus is a fantasy novel that tells the enchanting story of a magical competition between two young illusionists, Celia and Marco, who are bound by a powerful rivalry. The venue for their contest is a mysterious circus that only opens at night, filled with breathtaking wonders and intricate performances. As the two magicians push the limits of their abilities, they find themselves entwined in a deep and passionate connection that defies the very nature of their competition. The novel explores themes of love, destiny, and the blurred lines between reality and illusion.",
	}

	// Create a BookHelper instance
	bookHelper := helpers.NewBookHelper(book)

	// Load the HTML template
	tmpl, err := template.ParseFiles("./internal/views/templates/book.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template with the helper object
	tmpl.Execute(w, bookHelper)
}
