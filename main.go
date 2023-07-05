package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/partagile/lenslocked/controllers"
	"github.com/partagile/lenslocked/templates"
	"github.com/partagile/lenslocked/views"
)

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

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.Get("/chi/{userID}", chiRequestUrlParamHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Doh! 404 Not Found!", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
