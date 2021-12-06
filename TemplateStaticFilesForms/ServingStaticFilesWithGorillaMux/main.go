package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Person struct {
	Name string
	Id   string
}

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	// creating the http request multiplexer
	router := mux.NewRouter()
	// registering the function for parsing the template
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		ParsedTemplate, _ := template.ParseFiles("templates/first-template.html") // don't use "/templates/first-template.html"
		err := ParsedTemplate.Execute(rw, Person{Name: "Dinanath", Id: "1"})
		if err != nil {
			log.Fatal("Something went wrong while executing the Parsed template, ", err)
		}
	}).Methods("GET")
	/*
		http.StripPrefix("/static", http.FileServer(http.Dir("./static/")))
		This returns a handler that serves HTTP requests by removing /static from the request URL's path and invoking the file server.
		StripPrefix handles a request for a path that doesn't begin with prefix by replacing with an HTTP 404.
	*/
	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/")))) // don't use "/static/"
	// creating http server and catching the error
	if err := http.ListenAndServe(server_add, router); err != nil {
		log.Fatal("Something went wrong while starting the server, ", err)
	}
}
