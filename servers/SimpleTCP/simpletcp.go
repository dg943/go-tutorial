package main

import (
	"fmt"
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func handleConnRequests(conn net.Conn) {
	fmt.Println(conn)
}

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	listener, err := net.Listen(CONN_TYPE, server_add) /*
		function sign :-
		func Listen(network string, address string) (net.Listener, err)
		Here Listener is an interface which has the following functions
		1. Close() err
		2. Accept() (net.Conn, err)
		3. Addr()  net.Addr
	*/
	if err != nil {
		log.Fatal("Something went wrong while starting the tcp server : ", err)
	}
	defer listener.Close()
	fmt.Println("Listening on ", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Something went wrong while accepting the requests ", err)
		}
		fmt.Printf("Type of conn is : %T\n", conn)
		go handleConnRequests(conn) /*
			Handeling different tcp requests in different go routines
		*/
	}
}
