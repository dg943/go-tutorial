package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema" // this package fills a struct with form values
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type User struct {
	Username   string
	Password   string
	NotInclude string `schema:"-"` // this field is never set
}

func readForm(r *http.Request) *User {
	r.ParseForm()                                            // Parse the request body as a form and put the result into both r.PostForm and r.Form
	user := new(User)                                        // creating a new user or type *User i.e., it is a pointer
	decoder := schema.NewDecoder()                           // getting  new Decoder which we will be using to fill a user struct with form values
	if err := decoder.Decode(user, r.PostForm); err != nil { // popullating the user struct with the form
		log.Fatal("Something went wrong while decoding the form, ", err)
	}
	// Note : r.PostForm is only available after calling the r.ParseForm() function
	user.NotInclude = "set by me"
	return user
}

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			ParsedFiles, _ := template.ParseFiles("templates/login-form.html")
			if err := ParsedFiles.Execute(rw, nil); err != nil {
				log.Fatal("Something went wrong while executing the ParsedFiles, ", err)
			}
		} else {
			user := readForm(r)
			fmt.Fprintf(rw, "Hello "+user.Username+"!")
			fmt.Println(user)
		}
	})
	if err := http.ListenAndServe(server_add, nil); err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
