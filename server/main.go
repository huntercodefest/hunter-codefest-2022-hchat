package main

import (
	"fmt"
	"sync"
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
	fmt.Println("different")
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
