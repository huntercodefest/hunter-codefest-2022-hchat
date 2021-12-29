package main

import (
	"net"
)

const (
	TYPE = "tcp"
	HOST = "localhost"
	PORT = "4444"
)

func main() {
	startServer()
}

func startServer() {
	// Listen is server variable
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	// If server is not initialized correctly
	if err != nil {
		ExitOnErr(err)
	}
	// placeholder
	listen.Accept()

}
