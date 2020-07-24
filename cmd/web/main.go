/*
This file contains the main function responsible for settip up the app, it's dependencies and starting the server.
Author: Adarsh Kumar (iamadarshk.com)
Github: @iamadarshk
*/
package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.LUTC|log.Llongfile)
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.LUTC)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		ErrorLog: errorLog,
		Handler:  app.Routes(),
		Addr:     ":8080",
	}

	infoLog.Println("Starting server")
	errorLog.Fatal(srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem"))
}
