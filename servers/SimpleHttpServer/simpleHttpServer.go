package main

import (
	"fmt"
	"log"
	"net/http"
)

// defining the constants for creating the server address
const (
	CONN_PORT = "8080"
	CONN_HOST = "localhost"
)

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { /*
			function signature of HandleFunc :- func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
			HandleFunc registers the handle func with the given pattern in the DefaultServeMux
			In place of HandleFunc we could also use function Handle :- func Handle(patter string, handler Handler)
			Here Handler is an interface with method ServeHTTP :- func ServeHttp(ResponseWriter, *Request)
			Handle registers the handler with the given pattern in the DefaultServeMux
		*/
		fmt.Fprintf(rw, "Hello world!") /*
			function sign :- func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
			here http.ResponseWriter implements the interface io.Writer by providing the function definition of function :- Write(p []byte) (n int, err error)
			that is we are able to pass http.ResponseWriter as an input (polymorphism)
		*/
	})
	err := http.ListenAndServe(server_add, nil) /*
		function signature :- func ListenAndServe(add string, handler Handler) err
		ListenAndServe, listen on the TCP network address addr, and then Sever the handler to handle requests on incoming connections.

		In case of nil the handler is DefaultServeMux
		ListenAndServe always return a non nil error
	*/
	if err != nil {
		log.Fatal("Error staring the http server : ", err) /*
			function sign :- func (l *logger) Fatal(v ...interface{})
			It is equivalent to 1. calling l.Print() followed by os.Exit(1) non zero exit state code
			Also it can accept an infinite number of arguments of any type
		*/
	}
}
