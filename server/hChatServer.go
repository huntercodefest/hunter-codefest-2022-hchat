package main

import (
	"net"
)

const (
	TYPE = "tcp"
	PORT = "4444"
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
	room_num, username, _, err := ProcessInput(msgbuf)
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}
	// Validate username and room existence
	p_user, err := NewUser(conn, username, room_map)
	if err != nil {
		RespondWithErr(p_user, err)
		return
	}
	if _, ok := room_map[room_num]; !ok {
		room_map[room_num] = NewRoom(room_num, make([]*User, 1))
	}
	AddUserToRoom(p_user, room_map[room_num])
	//  seperate user into appropriate room
	return nil
}
