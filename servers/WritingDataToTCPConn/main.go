package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func handleRequest(conn net.Conn) {
	defer conn.Close()
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("Something's wrong")
	}
	fmt.Println(string(message))
	conn.Write([]byte("Hello from the server to the client\n"))
}

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	listener, err := net.Listen(CONN_TYPE, server_add)
	if err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
	defer listener.Close()
	fmt.Println("Listening on : ", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Something went wrong while accepting the tcp connection, ", err)
		}
		go handleRequest(conn)
	}
}

// command is : ncat localhost 8080
// output : Hello from the server to the client
