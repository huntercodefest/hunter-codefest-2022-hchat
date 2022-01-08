package main

import "fmt"

/*
Warning:
Please be careful when using Rooms to pass room pointer to avoid duplication
*/
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
		DelRoom((*p_room).room_num, room_map)
	}
}

// DelRoom() function deletes memory at passed in pointer for room
func DelRoom(room_num int, room_map map[int]*Room) {
	// room := Room{room_num: room_num, room_capacity: room_capacity, room_desc: room_desc, conn_users: conn_users}
	delete(room_map, room_num)
}

// TODO
// Rooms should be created and deleted if they contain users to prevent unneccessary bloat
// Room number and room description for all rooms should be stored on git and updated periodically, then searched through when creating new rooms

// CheckRoomHasConnection() function checks the passed in room to see if the user slice contains any users.

func CheckRoomHasConnection(p_room *Room) bool {
	return len((*p_room).conn_users) >= 1
}

func DistributeMessageToRoom(p_room *Room, message string) {
	fmt.Println("attempting distribute")
	fmt.Println("length of room: " + fmt.Sprint(len((*p_room).conn_users)))
	// loops through range of dereferenced room pointers user list
	for _, user := range (*p_room).conn_users {
		fmt.Println("in distribute loop")
		// attempts to send each user message
		err := RespondWithString(user, message)
		// if response fails passes user into remove user function
		if err != nil {
			RemoveUserFromRoom(user, p_room)
		}
	}
	fmt.Println("passed distribute loop")
}
