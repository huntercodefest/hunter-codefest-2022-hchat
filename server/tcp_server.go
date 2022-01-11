package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func startServer() {
	// Listen is server variable
	server, err := net.Listen("tcp", ":"+TCP_PORT)
	// If server is not initialized correctly
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("TCP Server open on port: " + TCP_PORT)
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
func processTCPConn(tcp_conn *net.Conn) (err error) {
	// func first needs to process message
	log.Println("TCP: client connected: " + (*tcp_conn).RemoteAddr().String())
	msgbuf, err := ReadTCPMessage(tcp_conn)
	if err != nil {
		return err
	}
	room_num, username, _, err := ProcessInput(msgbuf)
	if err != nil {
		(*tcp_conn).Write([]byte(err.Error()))
		return
	}
	// Validate username and room existence
	log.Println("created new tcp user")
	p_user, err := NewUser(tcp_conn, nil, username, ROOM_MAP)
	if err != nil {
		(*tcp_conn).Write([]byte(err.Error()))
		return
	}
	if _, ok := ROOM_MAP[room_num]; !ok {
		ROOM_MAP[room_num] = NewRoom(room_num, make([]*User, 0))
		log.Println("created new room #" + fmt.Sprint(room_num))
	}
	//  seperate user into appropriate room
	AddUserToRoom(p_user, ROOM_MAP[room_num])
	return nil
}
