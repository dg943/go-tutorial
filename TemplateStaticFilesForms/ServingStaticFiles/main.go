package main

import (
	"html/template"
	"log"
	"net/http"
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
	person := Person{Name: "Dinanath", Id: "1"}

	// There are these 2 lines which are extra

	fileserver := http.FileServer(http.Dir("./static")) // Here we are specifying the root directory from where the
	// http requests will be served, relative to the directory storing the go.mod file

	/*
		Here, we are striping the static portion of the url and passing the rest to the fileserver to serve the http request
	*/
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		ParsedTemplate, _ := template.ParseFiles("templates/first-template.html")
		err := ParsedTemplate.Execute(rw, person)
		if err != nil {
			log.Fatal("Something went wrong while executing the parsed template, ", err)
		}
	})

	err := http.ListenAndServe(server_add, nil)
	if err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
