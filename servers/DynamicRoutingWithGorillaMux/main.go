package main

import (
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
	router := mux.NewRouter() /*
		returns a http request multiplexer, which register the matching patterns, and the corresponding
		handlers
	*/
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello world!"))
	}).Methods("GET") // we can only access this path only, if the request header contains the GET request method
	router.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("It's a post request!"))
	}).Methods("POST") // we can only access this path only, if the request header contains the POST request method
	/*
		This is an example of dynamic routing, here we will be abel to access the variable after /hello/...
		with the help of key
	*/
	router.HandleFunc("/hello/{name}", func(rw http.ResponseWriter, r *http.Request) {
		path_variables := mux.Vars(r) /*
			function signature is :-
			func Vars(r *http.Request) map[string]string
			This function is returning a map where key and value both are values of type string
			requesting the request path variables
		*/
		name := path_variables["name"]  // accessing the value of key "name"
		rw.Write([]byte("Hi, " + name)) // writing to response writer
	}).Methods("GET", "PUT")
	err := http.ListenAndServe(server_add, router)
	if err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
