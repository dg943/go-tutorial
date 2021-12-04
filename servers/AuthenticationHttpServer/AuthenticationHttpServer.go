package main

import (
	"crypto/subtle" // Package subtle implements functions that are often useful in cryptographic code but require careful thought to use correctly.
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST      = "localhost"
	CONN_PORT      = "8080"
	ADMIN_USERNAME = "admin"
	ADMIN_PASSWORD = "admin"
)

/*
type HanlderFunc func(ResponseWriter, *Request)
The type HanlderFunc is an adapter to allow the use of ordinary functions as an http handler. If f is a function with the
appropriate signaure , HandlerFunc(f) is a hanlder that calls f.
*/
func Authenticate(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(username), []byte(ADMIN_USERNAME)) != 1 || subtle.ConstantTimeCompare([]byte(password), []byte(ADMIN_PASSWORD)) != 1 {
			/*
				function sign :- func ConstantTimeCompare(x, y []byte) int
				if x == y returns 1
				else return 0
			*/

			rw.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`) /*
				Header function returns the value of type Header which has the base type as a map of map[string][]string
				that is key := string
				and value := []string // a slice of strings
			*/
			rw.WriteHeader(401)                                                 // client side error
			rw.Write([]byte("You are not authorized to access the webpage.\n")) // Here we first have to use WriteHeader otherwise Write() function will set the status code to Ok or 200
			return
		}
		handler(rw, r)
	}
}

func helloworld1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world - 1!")
}

//func helloworld2(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello world!")
//}

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	http.HandleFunc("/", Authenticate(helloworld1, "Please enter your username and password"))
	err := http.ListenAndServe(server_add, nil)
	if err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
