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
	tables := []string{
		`CREATE TABLE IF NOT EXISTS users (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      email VARCHAR(50) UNIQUE NOT NULL,
      password VARCHAR(50) NOT NULL
  )`,
		`CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(255),
    location VARCHAR(255) NOT NULL,
	  datetime DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	)`,
		`CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		event_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(event_id) REFERENCES events(id) ON DELETE CASCADE
	)`,
	}

	for _, table := range tables {
		createTable(table)
	}
}

func createTable(script string) {
	_, exception := DB.Exec(script)
	if exception != nil {
		log.Fatal("Cannot create registrations table")
	}
}
