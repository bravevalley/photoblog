package app

import (
	"html/template"
	"net/http"
)

type Router struct {
	Mux *http.ServeMux
}

func (router *Router) Init(tpl *template.Template) {
	router.Mux = http.NewServeMux()
	router.serveRoute(tpl)
}

func (router *Router) serveRoute(tpl *template.Template) {
	h := &Handlers{
		Template: tpl,
		// Database: db,
		// Redis: rdb,
	}

	router.Mux.HandleFunc("/login", h.login)
	router.Mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// router.Mux.HandleFunc("/dashboard", h.dashboard)

}
