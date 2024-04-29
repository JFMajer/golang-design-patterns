package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowAbout(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	filename := fmt.Sprintf("%s.page.gohtml", page)

	if _, err := os.Stat("./templates/" + filename); os.IsNotExist(err) {
		app.renderNotFound(w)
		return
	}

	app.render(w, filename, nil)
}
