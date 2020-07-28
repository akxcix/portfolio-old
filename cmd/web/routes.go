// This file contains the routes and middleware integrations
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes for the web app
func (app *application) Routes() http.Handler {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(app.notFoundHandler)

	// Routes
	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/about", app.about).Methods("GET")

	// Static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer)).Methods("GET")

	return r
}
