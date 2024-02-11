package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	//Static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//View routes
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	//Middleware setup
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(mux)
}
