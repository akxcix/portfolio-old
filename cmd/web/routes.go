// This file contains the routes and middleware integrations
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes for the web app
func (app *application) Routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/{name}", app.greetings).Methods("GET")

	return r
}
