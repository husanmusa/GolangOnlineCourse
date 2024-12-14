package main

import (
	"fmt"

	"lesson18/postgres"

	"lesson18/storage"
)

func main() {
	db, err := postgres.ConnDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// msgStorage := storage.MessageStorage{db}
	msgStorage := storage.NewMessageStorage(db)

	messages, err := msgStorage.GetMessages(16456543, 16456543+10)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", messages)
}
