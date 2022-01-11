package main

import (
	"fmt"
	"net"
	"sync"
)

const (
	TYPE = "tcp"
	PORT = "5555"
	_LEN = 1024 // Max length of message
)

// Global scope map from room number to room
var ROOM_MAP map[int]*Room

func main() {
	// start tcp server and ws server as seperate goroutines
	// Use wait group to keep program running
	var wg sync.WaitGroup
	wg.Add(2)
	// TCP server code
	fmt.Println("different")
	go func() {
		fmt.Println("Do this first")
		startServer()
		fmt.Println("TCP SERVER FIN")
		wg.Done()
	}()
	// Websocket server code
	go func() {
		fmt.Println("Do this second")
		setupHTTPRoutes()
		fmt.Println("WS SERVER FIN")
		wg.Done()
	}()
	// servers should not hit this spot unless something goes wrong
	wg.Wait()
}

func startServer() {
	ROOM_MAP = make(map[int]*Room)
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
		go processInitialConnection(&conn)
	}
}

// Run as seperate Goroutine
func processInitialConnection(conn *net.Conn) (err error) {
	// func first needs to process message
	fmt.Println("New client connected: " + (*conn).RemoteAddr().String())
	msgbuf, err := ReadSingleMessage(conn)
	fmt.Println("passed read single message")
	if err != nil {
		return err
	}
	fmt.Println("hit process input")
	room_num, username, _, err := ProcessInput(msgbuf)
	fmt.Println("passed process input")
	if err != nil {
		(*conn).Write([]byte(err.Error()))
		return
	}
	// Validate username and room existence
	fmt.Println("created new user")
	p_user, err := NewUser(conn, username, ROOM_MAP)
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
