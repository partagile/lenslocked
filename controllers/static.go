package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is this the 1st question?",
			Answer:   "Yup!",
		},
		{
			Question: "I want 2!!",
			Answer:   "Thanks for your enthusiasm but... Not so fast! We've got that shipping problem :/",
		},
		{
			Question: "Yo! I'm not satisfied. How can I contact you guys?",
			Answer:   `Satisfaction guaranteed! If you find anything you don't like, give us a shout here: <a href="mailto:test@test.com">test@test.com</a>`,
		},
		{
			Question: "Where are you located?",
			Answer:   "The whole team is remote, dude!",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
