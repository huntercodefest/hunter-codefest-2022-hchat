package main

import (
	"errors"
	"net"
)

/*
Warning:
Please be careful when using Users to pass user pointer to avoid duplication
*/

// User struct
// Contains username and connection object
type User struct {
	user_conn net.Conn
	username  string
}

// Constructor
func NewUser(user_conn net.Conn, username string, allrooms map[int]*Room) (*User, error) {
	// Placeholder
	if ValidateUsername(username, allrooms) {
		return &(User{user_conn: user_conn, username: username}), nil
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
