package main

import (
	"net"
	"strings"

	"github.com/gorilla/websocket"
)

/*
Warning:
Please be careful when using Users to pass user pointer to avoid duplication
*/

// User struct
// Contains username and connection object

// TODO
// Add websocket.conn
// Add nils to correct spot in all new users of type
// Add conditional statement in write functions
// Or potentially create a user interface with tcpuser and wsuser structs

type User struct {
	tcp_conn *net.Conn
	ws_conn  *websocket.Conn
	username string
	room_num int
	readingOnLoop bool
}

// Constructor
// Passes in both TCP connection and Websocket connection
// One is initialized as nil to determine if connected user is tcp or ws client
func NewUser(tcp_conn *net.Conn, ws_conn *websocket.Conn, username string, room_num int) (*User, error) {
	return &(User{tcp_conn: tcp_conn, ws_conn: ws_conn, username: strings.TrimSpace(username), room_num: room_num, readingOnLoop: false}), nil
}

func DelUser(user User) {
	user = User{}
}

// Function to guarantee correct username format should be done on client side
// Pass in target username and an array of all the rooms
