package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Howdy!</h1>")
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

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/faq", faqHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Doh! 404 Not Found!", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
