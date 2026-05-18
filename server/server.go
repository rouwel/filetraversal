package server

import (
	"net/http"
	"html/template"
)

func HomeView(w http.ResponseWriter, r *http.Request) {
	tmplt, _ := template.ParseFiles("server/templates/index.html")
	err := tmplt.Execute(w, nil)
	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
	}
}