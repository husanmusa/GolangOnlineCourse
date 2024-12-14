package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "goonline"
)

type Message struct {
	Id        int
	UserId    int
	Message   string
	CreatedAt string
}

func main() {
	dbCon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	messages := []Message{}

	rows, err := db.Query("select id, user_id, message, created_at from message where id between $1 and $2", 16456543, 16456543+10)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var msg Message
		err = rows.Scan(&msg.Id, &msg.UserId, &msg.Message, &msg.CreatedAt)
		if err != nil {
			panic(err)
		}
		messages = append(messages, msg)
	}

	fmt.Printf("%+v\n", messages)
}
