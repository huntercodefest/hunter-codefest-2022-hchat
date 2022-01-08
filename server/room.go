package main

/*
Warning:
Please be careful when using Rooms to pass room pointer to avoid duplication
*/
// Room struct
// Contains array of connections and helper information
type Room struct {
	room_num      int    // Descrpiption of room for server
	room_desc     string // Description of room for client
	room_capacity int
	conn_users    [](*User) // Dynamic array of User pointers

}

// New Room takes in room number, room capacity, room desc, and pointer to empty slice of users (conn_users := make([]User, 0))
func NewRoom(room_num int, room_capacity int, room_desc string, conn_users *[]User) *Room {
	return &(Room{room_num: room_num, room_capacity: room_capacity, room_desc: room_desc, conn_users: conn_users})
}

// TODO
// When adding user to room create new goroutine to read for new incoming messages and send response back if new message is found
func AddUserToRoom(p_user *User, p_room *Room) (err error) {
	(*Room).conn_users.append()
	go ReadConnOnLoop(&((*p_user).user_con))
	return
}

// Search through array of users in room pointer
func removeUserFromRoom(p_user *User, p_room *Room) {
	i := &((*p_user).user_con) //assigning the position of the user that wants to be removed
	p_room.conn_users = append(p_room.conn_users[:i], p_room.conn_users[i+1:]...)
	hasUsers := CheckRoomHasConnection(p_room)
	if hasUsers == false {
		DelRoom()
	}
}

// DelRoom() function deletes memory at passed in pointer for room
func DelRoom(room_num int, room_capacity int, room_desc string, conn_users *[]User) *Room {
	// room := Room{room_num: room_num, room_capacity: room_capacity, room_desc: room_desc, conn_users: conn_users}
	p.room = Room{}
}

// TODO
// Rooms should be created and deleted if they contain users to prevent unneccessary bloat
// Room number and room description for all rooms should be stored on git and updated periodically, then searched through when creating new rooms

/*
CheckRoomHasConnection() function checks the passed in room to see if the
user slice contains any users.
*/

func CheckRoomHasConnection(p_room *Room) bool {
	return len((*p_room).conn_users) >= 1
}
