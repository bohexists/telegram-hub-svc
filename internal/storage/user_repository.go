package storage

import (
	"log"

	"github.com/bohexists/telegram-hub-svc/db"
)

// User represents the structure of the User table.
type User struct {
	ID        int64
	ChatID    int64
	FirstName string
	LastName  string
	Username  string
	CreatedAt string
}

// CreateUser inserts a new user record into the database.
func CreateUser(user User) error {
	query := `INSERT INTO users (chat_id, first_name, last_name, username, created_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING id`
	err := db.GetDB().QueryRow(query, user.ChatID, user.FirstName, user.LastName, user.Username).Scan(&user.ID)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	log.Printf("User created with ID: %d", user.ID)
	return nil
}
