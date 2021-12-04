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
	/*
		conn is also of type io.Reader, because net.Conn implements the function func Reader(b []byte) (int, err)
		and bufio.NewReader accepts value of type io.Reader and then we are calling ReadString function
		with delimeter as new line character, i.e., it will read the data from the incomming connection
		till it get a new line character
		it returns a string, and a value of type error
	*/
	message, err := bufio.NewReader(conn).ReadString('\n') // for bytes we have to use single quotes ''
	if err != nil {
		log.Fatal("something went wrong while reading the message from the client ", err)
	}
	fmt.Println("Recieve message is : ", string(message))
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
			log.Fatal("Something went wrong while accepting the connection request , ", err)
		}
		go handleRequest(conn)
	}
}
