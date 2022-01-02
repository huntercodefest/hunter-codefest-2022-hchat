package main

import (
	"errors"
	"fmt"
	"net"
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
func ReadConnOnLoop(p_usr_conn *net.Conn) (err error) {
	// infinite loop conditional
	conn_failed := false
	for !conn_failed {
		msgbuf, err := ReadSingleMessage(p_usr_conn)
		if err != nil {
			conn_failed = true
			break
			// TODO change to remove user connection from room if read fails (check err type to not be message failed to read)
		}
		room_num, username, message, err := ProcessInput(&msgbuf)
		// placeholder
		fmt.Println(room_num, username, message, err)
	}
	// placeholder
	return err
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
	#00000_username with no message can indicate entering and exiting chatroom
*/
func ProcessInput(msgbuf *[]byte) (room_number int, username string, message string, err error) {
	var input string
	// Buffer requires processing due to extra 0 bits after message
	for index, element := range *msgbuf {
		if element == 0 {
			// slice grabbing all elements of array before index
			input = string((*msgbuf)[:index])
		}
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
