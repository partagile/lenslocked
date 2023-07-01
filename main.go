package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Print("parsing template: %v", err)
		http.Error(w, "ERROR parsing template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Print("executing template: %v", err)
		http.Error(w, "ERROR executing template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>This is the contact page.</h1><p>To get in touch, email me at <a href=\"mailto:test@test.com\">test@test.com</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ page</h1>
	this is a multi-line string <br />
	gonna add some line breaks <br />
	<ul>
	<li><strong>Question 1</strong> - Yup!</li>
	<li><strong>Question 2</strong> - Not so fast</li>
	<li><strong>Question 3</strong> - Satisfaction guaranteed!</li>
	`)
}

// HTTP handler accessing the url routing parameters.
func chiRequestUrlParamHandler(w http.ResponseWriter, r *http.Request) {
	// fetch the url parameter `"userID"` from the request of a matching
	// routing pattern. An example routing pattern could be: /users/{userID}
	userID := chi.URLParam(r, "userID")
	fmt.Fprint(w, "hello: ", userID, " ... ")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/faq", faqHandler)
	r.Get("/contact", contactHandler)
	r.Get("/chi/{userID}", chiRequestUrlParamHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Doh! 404 Not Found!", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
