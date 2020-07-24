// This file contains the handlers which handle corresponding paths
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Method: GET, Path: "/"
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home"))
}

func (app *application) greetings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	str := fmt.Sprintf("Hello, %s", name)
	w.Write([]byte(str))
}
