package main

import "fmt"

// Global scope room map
var ROOM_MAP map[int]*Room

// Room struct
// Contains array of connections and helper information
type Room struct {
	room_num   int     // Description of room for server
	conn_users []*User // Dynamic array of User pointers

}

// New Room takes in room number, room capacity, room desc, and pointer to empty slice of users (conn_users := make([]User, 0))
func NewRoom(room_num int, conn_users []*User) *Room {
	return &(Room{room_num: room_num, conn_users: conn_users})
}

// TODO
// When adding user to room create new goroutine to read for new incoming messages and send response back if new message is found
func AddUserToRoom(p_user *User, p_room *Room) (err error) {
	fmt.Println("add user")
	(*p_room).conn_users = append((*p_room).conn_users, p_user)
	go ReadConnOnLoop(p_user)
	return
}

// Search through array of users in room pointer
func RemoveUserFromRoom(p_user *User, p_room *Room) {
	var user_index int
	conn_users := (*p_room).conn_users
	for index, element := range conn_users {
		if *element == *p_user {
			user_index = index
			break
		}
	}
	(*p_room).conn_users = append(conn_users[:user_index], conn_users[user_index+1:]...)
	hasUsers := CheckRoomHasConnection(p_room)
	if !hasUsers {
		DelRoom((*p_room).room_num)
	}
}

// DelRoom() function deletes room from globally scoped map of rooms
// Uses built in delete function to do so
func DelRoom(room_num int) {
	delete(ROOM_MAP, room_num)
}

// TODO
// Rooms should be created and deleted if they contain users to prevent unneccessary bloat
// Room number and room description for all rooms should be stored on git and updated periodically, then searched through when creating new rooms

// CheckRoomHasConnection() function checks the passed in room to see if the user slice contains any users.

func CheckRoomHasConnection(p_room *Room) bool {
	return len((*p_room).conn_users) >= 1
}

func DistributeMessageToRoom(p_room *Room, message string) {
	// loops through range of dereferenced room pointers user list
	for _, user := range (*p_room).conn_users {
		// attempts to send each user message
		err := RespondWithString(user, message)
		// if response fails passes user into remove user function
		if err != nil {
			RemoveUserFromRoom(user, p_room)
		}
	}
}
