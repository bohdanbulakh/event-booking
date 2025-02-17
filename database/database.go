package database

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

type Table map[string]string
type RowCallback func(row *sql.Row) error
type RowsCallback func(rows *sql.Rows) error

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		log.Fatal("Could not connect to database.")
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	tables := Table{
		"users": `CREATE TABLE IF NOT EXISTS users (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      email VARCHAR(50) UNIQUE NOT NULL,
      password VARCHAR(50) NOT NULL
  )`,
		"events": `CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(255),
    location VARCHAR(255) NOT NULL,
	  datetime DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	)`,
		"registrations": `CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		event_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(event_id) REFERENCES events(id) ON DELETE CASCADE
	)`,
	}

	for name, query := range tables {
		createTable(name, query)
	}
}

func createTable(name, query string) {
	_, exception := db.Exec(query)
	if exception != nil {
		log.Fatalf("Cannot create %s table:\n%s", name, exception)
	}
}

func Exec(query string, args ...any) (sql.Result, error) {
	statement, exception := db.Prepare(query)
	if exception != nil {
		return nil, exception
	}

	defer statement.Close()
	return statement.Exec(args...)
}

func Query(query string, callback RowsCallback, args ...any) (*sql.Rows, error) {
	rows, exception := db.Query(query, args...)
	if exception != nil {
		return nil, exception
	}

	defer rows.Close()
	if callback != nil {
		exception := callback(rows)
		if exception != nil {
			return nil, exception
		}
	}

	return rows, exception
}

func QueryRow(query string, callback RowCallback, args ...any) (*sql.Row, error) {
	row := db.QueryRow(query, args...)
	if row == nil {
		return nil, errors.New("no rows found")
	}

	if callback != nil {
		exception := callback(row)
		if exception != nil {
			return nil, exception
		}
	}

	return row, nil
}
