package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema" // this package fills a struct with form values
)

const (
	CONN_HOST              = "localhost"
	CONN_PORT              = "8080"
	USERNAME_ERROR_MESSAGE = "Please enter a valid username"
	PASSWORD_ERROR_MESSAGE = "Please enter a valid password"
	GENERIC_ERROR_MESSAGE  = "Validation error"
)

type User struct {
	Username string `valid:"alpha"`    // this valid tag is necessary for using the govalidator, here we are permiting only the alphabets as a username
	Password string `valid:"alphanum"` // Here, we are only promoting the use of alpha numeric characters for the password
}

func readForm(r *http.Request) *User {
	r.ParseForm()                                            // Parse the request body as a form and put the result into both r.PostForm and r.Form
	user := new(User)                                        // creating a new user or type *User i.e., it is a pointer
	decoder := schema.NewDecoder()                           // getting  new Decoder which we will be using to fill a user struct with form values
	if err := decoder.Decode(user, r.PostForm); err != nil { // popullating the user struct with the form
		log.Fatal("Something went wrong while decoding the form, ", err)
	}
	return user
}

func validateUser(user User) (bool, string) {
	valid, validationError := govalidator.ValidateStruct(user)
	if !valid {
		usernameError := govalidator.ErrorByField(validationError, "Username")
		passwordError := govalidator.ErrorByField(validationError, "Password")
		if usernameError != "" {
			log.Println("username validation error : ", usernameError)
			return valid, USERNAME_ERROR_MESSAGE
		}
		if passwordError != "" {
			log.Println("password validation error : ", passwordError)
			return valid, PASSWORD_ERROR_MESSAGE
		}
	}
	return valid, GENERIC_ERROR_MESSAGE
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
			valid, validationErrorMessage := validateUser(*user)
			if !valid {
				log.Println(validationErrorMessage)
				rw.Write([]byte(validationErrorMessage))
				return
			}
			fmt.Fprintf(rw, "Hello "+user.Username+"!")
			fmt.Println(user)
		}
	})
	if err := http.ListenAndServe(server_add, nil); err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
