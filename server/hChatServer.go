package main

import (
	"fmt"
	"net"
)

const (
	TYPE = "tcp"
	PORT = "5555"
	_LEN = 1024 // Max length of message
)

var room_map map[int]*Room

func main() {
	startServer()
}

func startServer() {
	room_map = make(map[int]*Room)
	// Listen is server variable
	server, err := net.Listen(TYPE, ":"+PORT)
	fmt.Println("Listening on port " + PORT)
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
	fmt.Println("New client connected: " + conn.RemoteAddr().String())
	msgbuf, err := ReadSingleMessage(&conn)
	fmt.Println("passed read single message")
	if err != nil {
		return err
	}
	fmt.Println("hit process input")
	room_num, username, _, err := ProcessInput(msgbuf)
	fmt.Println("passed process input")
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}
	// Validate username and room existence
	fmt.Println("created new user")
	p_user, err := NewUser(conn, username, room_map)
	if err != nil {
		RespondWithErr(p_user, err)
		return
	}
	if _, ok := room_map[room_num]; !ok {
		room_map[room_num] = NewRoom(room_num, make([]*User, 0))
		fmt.Println("created a new room at " + fmt.Sprint(room_num))
	}
	fmt.Println("hit add user")
	AddUserToRoom(p_user, room_map[room_num])
	fmt.Println("passed add user")
	//  seperate user into appropriate room
	return nil
}
