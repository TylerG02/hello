package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Create a multiplexer
	router := httprouter.New()
	// Create a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/create", app.Greeting)
	// ->home
	router.HandlerFunc(http.MethodGet, "/", app.Home)
	// ->about
	router.HandlerFunc(http.MethodGet, "/about", app.About)
	// Create our server
	router.HandlerFunc(http.MethodPost, "/create", app.MessageCreate)
	// Create our server
	return router
}
