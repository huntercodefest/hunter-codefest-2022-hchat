package main

import "github.com/gorilla/websocket"

// Suite of respond to client functions

// Overarching response function called by primary use respond functions
// Generally should not be used but is available in the case it is needed
func RespondToClient(p_user *User, msgbuf []byte) (err error) {
	if (*p_user).ws_conn == nil {
		// If user is connected via tcp
		_, err = (*(*p_user).tcp_conn).Write(msgbuf)
	} else {
		// If user is connected via WS
		err = (*(*p_user).ws_conn).WriteMessage(websocket.TextMessage, msgbuf)
	}
	return err
}

func RespondWithString(p_user *User, message string) (err error) {
	return RespondToClient(p_user, []byte(message))
}

func RespondWithErr(p_user *User, passed_err error) (err error) {
	return RespondToClient(p_user, []byte(passed_err.Error()))
}
