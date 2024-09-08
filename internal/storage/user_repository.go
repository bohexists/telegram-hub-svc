package storage

import (
	"database/sql"
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
func CreateUser(user User) (bool, error) {
	if UserExists(user.ChatID) {
		return false, nil // User already exists
	}
	query := `INSERT INTO users (chat_id, first_name, last_name, username, created_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING id`
	err := db.GetDB().QueryRow(query, user.ChatID, user.FirstName, user.LastName, user.Username).Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No rows were affected: %v", err)
			return false, err
		}
		log.Printf("Error creating user: %v", err)
		return false, err
	}
	log.Printf("User created with ID: %d", user.ID)
	return true, nil
}

// UserExists checks if a user exists in the database.
func UserExists(chatID int64) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE chat_id = $1)`
	err := db.GetDB().QueryRow(query, chatID).Scan(&exists)
	if err != nil {
		log.Printf("Error checking user existence: %v", err)
		return false
	}
	return exists
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers() ([]User, error) {
	users := []User{}
	rows, err := db.GetDB().Query(`SELECT id, chat_id, first_name, last_name, username, created_at FROM users`)
	if err != nil {
		log.Printf("Error getting users: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.ChatID, &u.FirstName, &u.LastName, &u.Username, &u.CreatedAt)
		if err != nil {
			log.Printf("Error scanning user: %v", err)
			continue
		}
		users = append(users, u)
	}
	return users, nil
}
