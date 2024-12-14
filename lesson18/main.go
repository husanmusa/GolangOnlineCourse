package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "goonline"
)

func main() {
	dbCon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	t := time.Now()

	var id int
	var user_id int
	var message string
	var created_at time.Time
	row := db.QueryRow("select id, user_id, message, created_at from message where id between 543254 and 543354")

	err = row.Scan(&id, &user_id, &message, &created_at)
	if err != nil {
		panic(err)
	}

	fmt.Println(id, user_id, message, created_at)
	fmt.Println(time.Since(t))

	// _, err = db.Exec("delete from message where id > 100000000")
	// if err != nil {
	// 	panic(err)
	// }

	// rows, err := res.RowsAffected()
	// if err != nil {
	// 	panic(err)
	// }
	// if rows == 0 {
	// 	fmt.Println("No rows affected, not found")
	// 	return
	// }
	// uuid := uuid.New()
	// message := "in Japan: Someone is waiting for you"
	// row = db.QueryRow("insert into message(user_id, message) values (1, $1) returning id", message)

	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// msg := struct {
	// 	Id        int
	// 	UserId    int
	// 	Message   string
	// 	CreatedAt time.Time
	// }{}

	// id = 7653464
	// err = db.QueryRow(`select id, user_id, message, created_at from message
	// where id = $1 or user_id = $2`, id, user_id).Scan(&msg.Id, &msg.UserId, &msg.Message, &msg.CreatedAt)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(time.Since(t), msg)
	// fmt.Printf("Id: %d\nUserId: %d\nMessage: %s\nCreatedAt: %s\n", msg.Id, msg.UserId, msg.Message, msg.CreatedAt)

	_, err = db.Exec("insert into message(user_id, message) values ($1, $2)",
		543253, "In Tashkent: Someone is waiting for you")
	if err != nil {
		panic(err)
	}
}
