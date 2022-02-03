package main

// Functions to open and query database of past messages

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type db_message struct {
	username string
	message  string
}

var db *sql.DB

func openDBConnection() {
	// Server is shared only on localhost and input is sanitized to prevent unauthorized access
	var err error
	// MySQL server hostname is stored on an unpushed file on host machine
	db, err = sql.Open("mysql", DB_CONN)

	if err != nil {
		panic(err.Error())
	}
}

func readDB(room_num int) (messages []db_message, err error) {
	limitQueries(room_num)
	results, err := db.Query("" +
	"SELECT username, message FROM messages WHERE " +
	 "room_num=? ORDER BY ID ASC", room_num)
	 if err != nil {
		panic(err)
		// return nil, err
	}
	for results.Next() {
		var msg db_message
		err := results.Scan(&msg.username, &msg.message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func getHighestID(room_num int) int {
	fmt.Println("Room num: ", room_num)
	result := db.QueryRow("SELECT ID FROM messages WHERE room_num=? ORDER BY ID DESC LIMIT 1", room_num)
	var ID int
	switch err := result.Scan(&ID); err{
	case sql.ErrNoRows:
			return 0
	case nil:
			return ID
	default:
		panic(err) 
	}
}
func writeToDB(room_num int, username string, message string) (err error) {
	max_id := getHighestID(room_num) + 1
	_, err = db.Exec("INSERT INTO messages VALUES(?, ?, ?, ?)", max_id, room_num, username, message)
	return err
}

// Store only latest 250 messages per room
func limitQueries(room_num int) {
	max_id := getHighestID(room_num)
	_, err := db.Exec("DELETE FROM messages WHERE ID<? AND room_num=?", max_id-250, room_num)
	if err != nil{
		panic(err)
	}
}
