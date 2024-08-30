package views

import (
	"html/template"
	"net/http"
)

func RenderUserView(w http.ResponseWriter, data interface{}) {
	// Render the template with the helper object
	tmpl, err := template.ParseFiles("./internal/views/templates/user-view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
