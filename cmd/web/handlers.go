package main

import (
	"errors"
	"fmt"
	"github.com/BurningBlueFox/letsgo/internal/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippetsModel.Latest(5)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{Snippets: snippets}
	app.render(w, r, http.StatusOK, "home.tmpl", data)

	app.logger.Info("served home")
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		app.notFound(w)
		return
	}

	snippet, err := app.snippetsModel.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := templateData{Snippet: snippet}
	app.render(w, r, http.StatusOK, "view.tmpl", data)

	app.logger.Info("served view", "ID", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O Snail"
	content := "O Snail\nClimb Mount Fuji\nBut slowly, slowly!\n\n- Kobayashi Issa"
	var expires uint = 7

	id, err := app.snippetsModel.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
	app.logger.Info("served create")
}
