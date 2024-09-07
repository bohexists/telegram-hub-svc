package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB initializes the PostgreSQL connectionß
func InitDB() {
	fmt.Println("Starting database initialization...")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	fmt.Println("Connection string:", connStr)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	fmt.Println("Database connection opened.")

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database: ", err)
	}

	log.Println("Connected to the PostgreSQL database!")
	fmt.Println("Connected to the PostgreSQL database!")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
