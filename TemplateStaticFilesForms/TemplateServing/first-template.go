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

func renderTemplate(rw http.ResponseWriter, r *http.Request) {
	p := Person{Name: "Dinanath", Id: "2xxyyzz"}
	/*
		function signature is
		func ParseFiles(fileNames ...string) (*template.Template, error)
	*/
	ParsedTemplate, _ := template.ParseFiles(`templates/first-template.html`)
	/*
		function signature is
		func Execute(wr io.Writer, data interface{}) error
		It is taking an input of type Writer, where it will be wirting the output to
	*/
	err := ParsedTemplate.Execute(rw, p)
	if err != nil {
		log.Fatal("Something went wrong while executing the template, ", err)
	}
}

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(server_add, nil)
	if err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
