package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
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

func ReadConnForNewMessage(user User) (message string) {
	// infinite loop conditional
	conn_failed := false
	for !conn_failed {
		// Buffer slice to contain sent message
		msgbuf := make([]byte, _LEN)
		_, err := user.user_con.Read(msgbuf)
		if err != nil {
			conn_failed = true
		}
		// Buffer requires processing due to extra 0 bits after message
		var input string
		for index, element := range msgbuf {
			if element == 0 {
				// slice grabbing all elements of array before index
				input = string(msgbuf[:index])
				break
			}
		}
		room_num, username, message, err := ProcessInput(input)
		// placeholder
		fmt.Println(room_num, username, message, err)
	}
	// placeholder
	return
}

// func to process message
/*
	Message Syntax:
	#00000 - To indicate room
	_username - To indicate user
	:message - To indicate message
	<----------------------------------------------->
	The full syntax should look something like this:
	#00000_username:message
	#00000_username with no message can indicate entering and exiting chatroom
*/
func ProcessInput(input string) (room_number int, username string, message string, err error) {
	// if input does not contain #, _, :
	if input[0] != '#' || !strings.ContainsRune(input, '_') || !strings.ContainsRune(input, ':') {
		return 0, "", "", errors.New("wrong format, please consult github readme")
	}
	inlen := len(input)
	for i := 1; i < inlen; i++ {
		if input[i] == '_' {
			room_number, _ = strconv.Atoi(input[1:i])
			for j := i; j < inlen; j++ {
				if input[j] == ':' {
					username = input[i+1 : j]
					message = input[j:]
					return
				}
			}
		}
	}
	// placeholder
	return 0, "", "", errors.New("wrong format, please consult github readme")
}
