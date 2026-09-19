package handler

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// Embed the entire templates directory
//
//go:embed templates
var templateFS embed.FS

// Handler is the entry point Vercel calls for incoming web traffic
func Handler(w http.ResponseWriter, r *http.Request) {
	// Normalize request path
	path := strings.TrimSuffix(r.URL.Path, "/")

	// Route definitions
	switch path {
	case "", "/", "/api", "/api/index", "/api/index.go":
		home(w, r)
	case "/about", "/api/about":
		about(w, r)
	default:
		http.NotFound(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Parse template directly from embedded memory
	t, err := template.ParseFS(templateFS, "templates/"+tmpl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Template Error: %v\nCheck that api/templates/%s exists.", err, tmpl)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Render Error: %v", err)
	}
}

// package handler

// import (
// 	"embed"
// 	"html/template"
// 	"net/http"
// )

// // Embed the template directory so all html files are available at runtime.
// // Using the directory pattern is more robust than matching a glob directly.
// //
// //go:embed templates
// var templateFS embed.FS

// // Handler is the entry point Vercel calls for incoming web traffic
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	// Standardize routing path manually since we don't have a main loop routing for us
// 	path := r.URL.Path

// 	// Route definitions matching your main.go setup
// 	if path == "/" {
// 		home(w, r)
// 		return
// 	} else if path == "/about" {
// 		about(w, r)
// 		return
// 	}

// 	// fallback 404 if route doesn't match
// 	http.NotFound(w, r)
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "home.html")
// }

// func about(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "about.html")
// }

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	// Vercel maps your root files into the runtime path execution.
// 	// We read directly from the bundled templates folder.
// 	t, err := template.ParseFS(templateFS, "templates/"+tmpl)
// 	if err != nil {
// 		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	t.Execute(w, nil)
// }
