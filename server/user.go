package main

import (
	"errors"
	"net"
)

// User struct
// Contains username and connection object
type User struct {
	user_con net.Conn
	username string
}

// Constructor
func NewUser(user_con net.Conn, username string) (*User, error) {
	if ValidateUsername(username) {
		return &(User{user_con: user_con, username: username}), nil
	}
	return nil, errors.New("error: failed username check")
}

//  Function to guarantee correct username format, may be better to do on client side
func ValidateUsername(username string) (valid bool) {
	return
}
