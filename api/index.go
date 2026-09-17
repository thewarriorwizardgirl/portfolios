package handler

import (
	"html/template"
	"net/http"
)

// Handler is the entry point Vercel calls for incoming web traffic
func Handler(w http.ResponseWriter, r *http.Request) {
	// Standardize routing path manually since we don't have a main loop routing for us
	path := r.URL.Path

	// Route definitions matching your main.go setup
	if path == "/" {
		home(w, r)
		return
	} else if path == "/about" {
		about(w, r)
		return
	}

	// fallback 404 if route doesn't match
	http.NotFound(w, r)
}

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Vercel maps your root files into the runtime path execution.
	// We read directly from the bundled templates folder.
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
