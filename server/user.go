package main

import (
	"errors"
	"net"

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
}

// Constructor
// Passes in both TCP connection and Websocket connection
// One is initialized as nil to determine if connected user is tcp or ws client
func NewUser(tcp_conn *net.Conn, ws_conn *websocket.Conn, username string, allrooms map[int]*Room) (*User, error) {
	if ValidateUsername(username, allrooms) {
		return &(User{tcp_conn: tcp_conn, ws_conn: ws_conn, username: username}), nil
	}
	return nil, errors.New("error: failed username check")
}

func DelUser(user User) {
	user = User{}
}

// Function to guarantee correct username format should be done on client side
// Pass in target username and an array of all the rooms
func ValidateUsername(username string, p_room_arr map[int]*Room) (valid bool) {
	// Loops through connected Users & checks for duplicate username
	for i := range p_room_arr {
		for j := range p_room_arr[i].conn_users {
			if username == p_room_arr[i].conn_users[j].username {
				return false
			}
		}
	}
	return true
}
