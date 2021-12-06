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

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	router := mux.NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
		if err := parsedTemplate.Execute(rw, nil); err != nil {
			log.Fatal("Something went wrong while executing the parsed template, ", err)
		}
	})
	if err := http.ListenAndServe(server_add, router); err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
