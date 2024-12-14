package storage

import (
	"database/sql"

	"lesson18/models"
)

type MessageStorage struct {
	DB *sql.DB
}

func NewMessageStorage(db *sql.DB) *MessageStorage {
	return &MessageStorage{DB: db}
}

func (ms *MessageStorage) CreateMessage(message models.Message) error {
	_, err := ms.DB.Exec("insert into message(user_id, message) values ($1, $2)", message.UserId, message.Message)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MessageStorage) GetMessage(id int) (models.Message, error) {
	row := ms.DB.QueryRow("select id, user_id, message, created_at from message where id = $1", id)

	var msg models.Message
	err := row.Scan(&msg.Id, &msg.UserId, &msg.Message, &msg.CreatedAt)
	if err != nil {
		return models.Message{}, err
	}

	return msg, nil
}

func (ms *MessageStorage) GetMessages(from, to int) ([]models.Message, error) {
	messages := []models.Message{}

	rows, err := ms.DB.Query("select id, user_id, message, created_at from message where id between $1 and $2", from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var msg models.Message
		err = rows.Scan(&msg.Id, &msg.UserId, &msg.Message, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

// Update, Delete - CRUD - create, read, update, delete
