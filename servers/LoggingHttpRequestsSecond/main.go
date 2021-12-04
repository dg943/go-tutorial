/*
This is a sorter version of the logging programme
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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
		rw.Write([]byte("Hello world!!"))
	})
	router.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("This is the post request!!"))
	})
	router.HandleFunc("/hello/{name}", func(rw http.ResponseWriter, r *http.Request) {
		path_variable := mux.Vars(r)
		name := path_variable["name"]
		rw.Write([]byte("Hello! " + name))
	})

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	fmt.Println(logFile)
	if err != nil {
		log.Fatal("Error in starting or opening the file, ", err)
	}
	// since, Router has the function called ServeHTTP therefore it is also of type http.Handler
	loggedRouter := handlers.CombinedLoggingHandler(logFile, router)
	/*
		function signature of the http.ListenAndServe
		func ListenAndServe(add string, handler http.Handler) error
	*/
	err = http.ListenAndServe(server_add, loggedRouter)
	if err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
