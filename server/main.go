package main

import (
	"fmt"
	"sync"
)

const (
	HTTP_PORT = "8080"
	TCP_PORT  = "5555"
	MSG_LEN   = 1024
)

// Main function
func main() {
	// Make globally scoped ROOM_MAP variable
	ROOM_MAP = make(map[int]*Room)

	// start tcp server and ws server as seperate goroutines
	// Use wait group to keep program running
	var wg sync.WaitGroup
	wg.Add(2)
	// TCP server code
	go func() {
		startServer()
		fmt.Println("TCP SERVER FIN")
		wg.Done()
	}()
	// Websocket server code
	go func() {
		setupHTTPRoutes()
		fmt.Println("WS SERVER FIN")
		wg.Done()
	}()
	// servers should not hit this spot unless something goes wrong
	wg.Wait()
}
