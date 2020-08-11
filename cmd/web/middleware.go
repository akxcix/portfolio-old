package main

import (
	"net/http"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s\t%s\t%s\t%s\t%s", r.Proto, r.RemoteAddr, r.Method, r.URL, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}
