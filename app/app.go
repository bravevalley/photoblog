package app

import (
	"html/template"
	"net/http"
)

type App struct {
	Router *Router
	Server *http.Server
}

func (app *App) Init(tpl *template.Template) {
	app.Router = &Router{}
	app.Router.Init(tpl)
}

func (app *App) Run() {
	app.Server = &http.Server{
		Addr:    "8080",
		Handler: app.Router.Mux,
	}

	app.Server.ListenAndServe()
}
