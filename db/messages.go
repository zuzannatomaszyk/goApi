package db

import (
	"github.com/zuzannatomaszyk/goApi/models"
)

func (db Database) GetMessages() (*models.MessagesList, error) {
	//query execution
	list := &models.MessagesList{}
	rows, err := db.Conn.Query("SELECT createdAt, username,  messageContent FROM messages ORDER BY createdAt DESC fetch first 100 rows only")
	if err != nil {
		return list, err
	}
	//add every row returned by query to messages array
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.Timestamp, &message.User, &message.Text)
		if err != nil {
			return list, err
		}
		list.Messages = append(list.Messages, message)
	}
	return list, nil
}
func (db Database) AddMessage(message *models.Message) error {
	//first insert username to users table if user doesn't already exist in database
	query := `INSERT INTO users (username) VALUES ($1) ON CONFLICT DO NOTHING`
	db.Conn.QueryRow(query, message.User)

	//then insert message to messages table 
	var createdAt string
	query = `INSERT INTO messages (username,  messageContent) VALUES ($1, $2) RETURNING createdAt`
	err := db.Conn.QueryRow(query, message.User, message.Text).Scan(&createdAt)
	if err != nil {
		return err
	}
	message.Timestamp = createdAt
	return nil
}