/*
Logging http request is always a good idea while troubleshooting a web application, so it's a good
idea to log a request/response with a proper message and logging level. Go provides package "log", which
can help us to implement logging in an application. However we will be using the gorilla logging handlers
to implement it because the library offers more features such as logging in Apache Combined Log Format and
Apache Common Log Format, Which are not yet supported by the golang log package
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

type GetRequestHandler struct{}

func (h GetRequestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello world!!"))
}

type PostRequestHandler struct{}

func (h PostRequestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("This is the post request!!"))
}

type PathVariableHandler struct{}

func (h PathVariableHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path_variables := mux.Vars(r)
	name := path_variables["name"]
	rw.Write([]byte("Hi! " + name))
}

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	router := mux.NewRouter() /*
		type of router is *mux.Router, it also has a function ServeHTTP, therefore it is also of type http.Handler
	*/

	/*
		handlers.logingHandler function signature
		func loginHandler(io.Writer, http.Handler) http.Handler
	*/
	router.Handle("/", handlers.LoggingHandler(os.Stdout, new(GetRequestHandler))).Methods("GET") // this is going to write the logs on standtout output stream with apache common log format (clf)

	// creating a file called server.log with pointer as logFile, with 3 tags, and permission 0666
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	fmt.Println(logFile)
	if err != nil {
		log.Fatal("Error in starting or opening the file, ", err)
	}

	// registering another handler with the path "/post" and the format is still the clf but the logs will be written on the file server.log
	router.Handle("/post", handlers.LoggingHandler(logFile, new(PostRequestHandler))).Methods("POST")

	// here again we are registering another handler with the dynamic url "/hello/{name}" and the format is apache combined log format
	router.Handle("/hello/{name}", handlers.CombinedLoggingHandler(logFile, new(PathVariableHandler))).Methods("GET")

	err = http.ListenAndServe(server_add, router)
	if err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
