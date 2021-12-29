package main

// Room struct
// Contains array of connections and helper information
type Room struct {
	room_num      int    // Descrpiption of room for server
	room_desc     string // Description of room for client
	room_capacity int
	conn_users    *[]User // Dynamic array of User struct

}

// New Room takes in room number, room capacity, room desc, and pointer to empty slice of users (conn_users := make([]User, 0))
func NewRoom(room_num int, room_capacity int, room_desc string, conn_users *[]User) *Room {
	return &(Room{room_num: room_num, room_capacity: room_capacity, room_desc: room_desc, conn_users: conn_users})
}

// TODO
func AddUser(p_usr *User) (err error) {
	return
}

// TODO
// Rooms should be created and deleted if they contain users to prevent unneccessary bloat
// Room number and room description for all rooms should be stored on git and updated periodically, then searched through when creating new rooms
