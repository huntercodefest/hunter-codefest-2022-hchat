package main

import (
	"fmt"
	"net"
)

const (
	TYPE = "tcp"
	HOST = "localhost"
	PORT = "4444"
	_LEN = 1024 // Max length of message
)

func main() {
	startServer()
}

func startServer() {
	// Listen is server variable
	server, err := net.Listen(TYPE, HOST+":"+PORT)
	// If server is not initialized correctly
	if err != nil {
		ExitOnErr(err)
	}
	// Main loop to handle accepting connections
	for {
		conn, err := server.Accept()
		// If connection fails
		if err != nil {
			continue
		}
		// Seperates connection handle function from loop to allow main server to continue allowing new connections
		go processInitialConnection(conn)
	}
}

// Run as seperate Goroutine
func processInitialConnection(conn net.Conn) (err error) {
	// func first needs to process message
	msgbuf, err := ReadSingleMessage(&conn)
	if err != nil {
		return err
	}
	room_num, username, _, err := ProcessInput(&msgbuf)
	fmt.Println(room_num, username, err)
	// then validate username and room connection
	// then seperate user into appropriate room
	// then open new goroutine to read for new incoming messages from user to distribute to room
	return nil
}
