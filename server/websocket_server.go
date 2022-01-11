package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
This file is a very basic first version of a websocket server
It is a one to one clone of a video I found on https://tutorialedge.net/golang/go-websocket-tutorial/
Down the line I will adapt it to fit our needs
*/

var upgrader = websocket.Upgrader{
	ReadBufferSize:  _LEN,
	WriteBufferSize: _LEN,
}

func reader(conn *websocket.Conn) {
	// TODO
	// Find out if ReadMessage and WriteMessage are interchangable with Read and Write

	// If theres no function to write to either use switch to check if type is of websocket or of tcp connection
	// After that forget about this function entirely and adapt to processRoutine function
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client succesfully connected...")
	reader(ws)

}

func setupHTTPRoutes() {
	fmt.Println("Setting up routes")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
