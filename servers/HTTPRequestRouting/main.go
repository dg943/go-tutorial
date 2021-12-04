package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	// all these handlers will get registered with the DefaultServeMux
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello world!!")
	})
	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "login page!")
	})
	http.HandleFunc("/logout", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "logout page!")
	})
	// nil significe we are using the DefaultServeMux
	err := http.ListenAndServe(server_add, nil)
	if err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
