//GZIP compression sends the response from the server to the client in the .gzip format rather than sending the plane
//text format and it is always a good practice to send the compressed response if the client/browser supports it.

//By sending the compressed response we are saving the network bandwidth and reducing the download time eventually rendering
//the page faster
//what happens is that the request header sent by the browser contains the field that tells the server that it accepts compressed
//content (.gzip and .deflate) and if the server has the capability to send the response in compressed format then sent it.
//If the server supports compression then it will send the data in compressed format and set the field "Content-Encoding: gzip" as a respone header
//otherwise it sends a plan response back to the client , which clearly means asking for compresed  response is only a request
//by the browser and not a demand.
//
//we will be using Gorilla's handler package to implement it

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT // creating the server address
	mux := http.NewServeMux()                 /*
		Creating new Http request multiplexer for mapping the incoming url to the correct pattern and calling the corresponding handler
		It returns a pointer of the type *ServeMux (type is ServeMux which also implements the function SeverHttp, therefore it
		is also of type http.Handler)
	*/
	mux.HandleFunc("/", helloworld)                                       // assigning the function helloworld to the path "/"
	err := http.ListenAndServe(server_add, handlers.CompressHandler(mux)) /*
		Here, http.ListenAndServe has 2 parameters
		1. add string, which is the server addres, in our case it localhost:8080
		2. handler http.Hanlder , since the base type of ServeMux is http.Handler so we can use converson
	*/
	if err != nil { // in the end we are just catching any error while starting the server
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}

// output
// In the output, you will see the "Content-Encoding: gzip" filed in the response header will be there
// which is an indicator that the response sent by the server is in gzip format
