// This file contains the main function responsible for settip up the app, it's dependencies and starting the server.
// Author: Adarsh Kumar (iamadarshk.com)
// Github: @iamadarshk
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	tls := flag.Bool("tls", false, "Specifies wether to use TLS or not.")
	addr := flag.String("addr", ":443", "Address the server should listen at.")
	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.LUTC|log.Llongfile)
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.LUTC)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		ErrorLog: errorLog,
		Handler:  app.Routes(),
		Addr:     *addr,
	}

	infoLog.Printf("Starting server with TLS=%t, listening at %s", *tls, *addr)
	if *tls {
		errorLog.Fatal(srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem"))
	} else {
		errorLog.Fatal(srv.ListenAndServe())
	}
}
