package main

import (
	"log"
	"net/http"
	handler "portfolios/api"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// func main() {
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
// 	// })
// 	// Serve static files from the static directory
// 	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public/"))))

// 	http.HandleFunc("/", home)
// 	http.HandleFunc("/about", about)

// 	http.ListenAndServe(":8080", nil)
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	//fmt.Fprintf(w, "Hello, you've requested custom template home: %s\n", r.URL.Path)
// 	renderTemplate(w, "home.html")

// }

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	t, err := template.ParseFiles("templates/" + tmpl)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	//Execute the template to the writer because there is no error
// 	t.Execute(w, nil)

// }

// func about(w http.ResponseWriter, r *http.Request) {
// 	//fmt.Fprintf(w, "Hello, you've requested custom template about: %s\n", r.URL.Path)
// 	renderTemplate(w, "about.html")
// }
