package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

// Basic read from connection functionality
func ReadTCPMessage(p_usr_conn *net.Conn) ([]byte, error) {
	// Buffer slice to contain sent message
	msgbuf := make([]byte, MSG_LEN)
	_, err := (*p_usr_conn).Read(msgbuf)
	return msgbuf, err
}

func ReadWSMessage(p_ws_conn *websocket.Conn) ([]byte, error) {
	_, msgbuf, err := (*p_ws_conn).ReadMessage()
	return msgbuf, err
}

// function should run in gorutine on each new user to read its new messages
func ReadConnOnLoop(p_user *User) (err error) {
	if ((*p_user).readingOnLoop){
		return
	}
	(*p_user).readingOnLoop = true
	log.Println("starting read conn on loop")
	// Send connection message
	RespondWithString(p_user, "Succesfully connected to room\n")
	// infinite read loop
	tcp_check := ((*p_user).tcp_conn != nil)
	for {
		var msgbuf []byte
		if tcp_check {
			msgbuf, err = ReadTCPMessage((*p_user).tcp_conn)
		} else {
			msgbuf, err = ReadWSMessage((*p_user).ws_conn)
		}
		if err != nil {
			return err
		}
		room_num, username, message, err := ProcessInput(msgbuf)
		fmt.Println(room_num)
		if err != nil {
			return err
		}
		if username != (*p_user).username{
			(*p_user).username = username
		}
		if room_num != (*p_user).room_num{
			RemoveUserFromRoom(p_user, (*p_user).room_num)
			(*p_user).room_num = room_num
			AddUserToRoom(p_user, room_num)
			log.Println("Moved user to room: " + fmt.Sprint(room_num))
		}
		log.Println("Received msg: " + string(msgbuf) + " from user: " + username)
		log.Println("Received msg: " + message + " from user: " + username)
		DistributeMessageToRoom(ROOM_MAP[room_num], username+":"+message)
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
		return -1, "", "", errors.New("wrong format, please consult github readme")
	}
	inlen := len(input)
	for i := 1; i < inlen; i++ {
		if input[i] == '_' {
			room_number, _ = strconv.Atoi(input[1:i])
			fmt.Println(room_number)
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
