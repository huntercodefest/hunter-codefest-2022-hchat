package main

// Functions to open and query database of past messages

import (
	"database/sql"

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
	db, err = sql.Open("mysql", DB_CONN)

	if err != nil {
		panic(err.Error())
	}
	// Close DB connection if server fails or closes
	// defer db.Close()
}

func readDB(room_num int) (messages []db_message, err error) {
	limitQueries(room_num)
	results, err := db.Query("" +
	"SELECT username, message FROM messages WHERE " +
	 "room_num=? ORDER BY ID DESC", room_num)
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

func getHighestID() int {
	result, err := db.Query("SELECT ID from messages ORDER BY ID DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	var ID int
	result.Scan(&ID)
	return ID
}
func writeToDB(room_num int, username string, message string) (err error) {
	max_id := getHighestID()
	_, err = db.Exec("INSERT INTO messages VALUES(?, ?, ?, ?)", max_id + 1, room_num, username, message)
	return err
}

// Store only latest 250 messages
func limitQueries(room_num int) {
	max_id := getHighestID()
	_, err := db.Exec("DELETE FROM messages WHERE ID<? AND room_num=?", max_id-250, room_num)
	if err != nil{
		panic(err)
	}
}
