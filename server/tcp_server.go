package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	TYPE = "tcp"
	PORT = "5555"
	_LEN = 1024 // Max length of message
)

func startServer() {
	// Listen is server variable
	server, err := net.Listen(TYPE, ":"+PORT)
	fmt.Println("Listening on port " + PORT)
	// If server is not initialized correctly
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// Main loop to handle accepting connections
	for {
		conn, err := server.Accept()
		// If connection fails
		if err != nil {
			continue
		}
		// Seperates connection handle function from loop to allow main server to continue allowing new connections
		go processTCPConn(&conn)
	}
}

// Run as seperate Goroutine
func processTCPConn(conn *net.Conn) (err error) {
	// func first needs to process message
	log.Println("TCP: client connected: " + (*conn).RemoteAddr().String())
	msgbuf, err := ReadTCPMessage(conn)
	if err != nil {
		return err
	}
	room_num, username, _, err := ProcessInput(msgbuf)
	if err != nil {
		(*conn).Write([]byte(err.Error()))
		return
	}
	// Validate username and room existence
	fmt.Println("created new user")
	p_user, err := NewUser(conn, nil, username, ROOM_MAP)
	if err != nil {
		(*conn).Write([]byte(err.Error()))
		return
	}
	if _, ok := ROOM_MAP[room_num]; !ok {
		ROOM_MAP[room_num] = NewRoom(room_num, make([]*User, 0))
		fmt.Println("created a new room at " + fmt.Sprint(room_num))
	}
	fmt.Println("hit add user")
	AddUserToRoom(p_user, ROOM_MAP[room_num])
	fmt.Println("passed add user")
	//  seperate user into appropriate room
	return nil
}
