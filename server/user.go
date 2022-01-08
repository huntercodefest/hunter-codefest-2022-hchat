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
	user_con  net.Conn
	username  string
	user_room *Room
}

// Constructor
func NewUser(user_con net.Conn, username string, room *Room) (*User, error) {
	// Placeholder
	if ValidateUsername(username, nil) {
		return &(User{user_con: user_con, username: username, user_room: room}), nil
	}
	return nil, errors.New("error: failed username check")
}

// Function to guarantee correct username format should be done on client side
// Server side validation will search through all connected clients and check none have the same username
func ValidateUsername(username string, p_room_arr *[]Room) (valid bool) {
	return true
}
