package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// Basic read from connection functionality
func ReadSingleMessage(p_usr_conn *net.Conn) (input []byte, err error) {
	// Buffer slice to contain sent message
	msgbuf := make([]byte, _LEN)
	_, err = (*p_usr_conn).Read(msgbuf)
	if err != nil {
		return nil, err
	}
	return msgbuf, nil
}

// function should run in gorutine on each new user to read its new messages
func ReadConnOnLoop(p_user *User) (err error) {
	// Send connection message
	RespondWithString(p_user, "Succesfully connected to room\n")
	// infinite read loop
	for {
		msgbuf, err := ReadSingleMessage(&(*p_user).user_conn)
		if err != nil {
			return err
		}
		room_num, username, message, err := ProcessInput(msgbuf)
		if err != nil {
			return err
		}
		fmt.Println("distribute message to room hit")
		DistributeMessageToRoom(ROOM_MAP[room_num], username+":"+message)
		fmt.Println("passed distribute message to room")
	}
}

// func to process message format
/*
	Message Syntax:
	#00000 - To indicate room
	_username - To indicate user
	:message - To indicate message
	<----------------------------------------------->
	The full syntax should look something like this:
	#00000_username:message
	#00000_username: with no message can indicate entering and exiting chatroom
*/
func ProcessInput(msgbuf []byte) (room_number int, username string, message string, err error) {
	var input string
	// Buffer requires processing due to extra 0 bits after message
	foundBuffEnd := false
	for index, element := range msgbuf {
		if element == 0 {
			// slice grabbing all elements of array before index
			input = string((msgbuf)[:index])
			foundBuffEnd = true
		}
	}
	// If passed in message is max length and no zeros are found
	if !foundBuffEnd {
		input = string(msgbuf)
	}
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
					message = input[j+1:]
					return
				}
			}
		}
	}
	// placeholder
	return 0, "", "", errors.New("wrong format, please consult github readme")
}

// Suite of respond to client functions

// Overarching response function called by primary use respond functions
// Generally should not be used but is available in the case it is needed
func RespondToClient(p_user *User, msgbuf []byte) (err error) {
	fmt.Println("hit respond to client")
	conn := (*p_user).user_conn
	_, err = conn.Write(msgbuf)
	return err
}

func RespondWithString(p_user *User, message string) (err error) {
	return RespondToClient(p_user, []byte(message))
}

func RespondWithErr(p_user *User, passed_err error) (err error) {
	return RespondToClient(p_user, []byte(passed_err.Error()))
}

// For convienience sake only
func ExitOnErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}
