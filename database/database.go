package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		log.Fatal("Could not connect to database.")
	}

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUsersTableQuery := `
  CREATE TABLE IF NOT EXISTS users (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      email VARCHAR(50) UNIQUE NOT NULL,
      password VARCHAR(50) NOT NULL
  )`
	_, exception := DB.Exec(createUsersTableQuery)
	if exception != nil {
		log.Fatal("Cannot create users table")
	}

	createEventsTableSQL := `
	CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(255),
    location VARCHAR(255) NOT NULL,
	  datetime DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	)`
	_, exception = DB.Exec(createEventsTableSQL)
	if exception != nil {
		log.Fatal("Cannot create events table")
	}
}
