package models

import (
	"database/sql"
	"event-booking/database"
	"time"
)

type scannable interface {
	Scan(dest ...any) error
}

type Event struct {
	Id          int64
	Name        string    `binding:"required" json:"name"`
	Description string    `json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"datetime"`
	UserId      int64
}

func (event *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	result, exception := database.Exec(query,
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserId)

	if exception != nil {
		return exception
	}

	id, exception := result.LastInsertId()
	if exception != nil {
		return exception
	}

	event.Id = id
	return nil
}

func getEvent(row scannable) (Event, error) {
	var event Event
	exception := row.Scan(
		&event.Id,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId,
	)

	return event, exception
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	var events []Event

	_, exception := database.Query(query, func(rows *sql.Rows) error {
		for rows.Next() {
			event, exception := getEvent(rows)
			if exception != nil {
				return exception
			}

			events = append(events, event)
		}
		return nil
	})

	return events, exception
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	var event Event

	_, exception := database.QueryRow(query, func(row *sql.Row) error {
		var exception error
		event, exception = getEvent(row)

		return exception
	}, id)

	if exception != nil {
		return nil, exception
	}

	return &event, exception
}

func (event *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, datetime = ?, user_id = ?
	WHERE id = ?`

	_, exception := database.Exec(
		query,
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserId,
		event.Id,
	)

	return exception
}

func (event *Event) Delete() error {
	query := "DELETE FROM events where id = ?"
	_, exception := database.Exec(query, event.Id)

	return exception
}

func (event *Event) Register(userId int64) error {
	query := "INSERT INTO registrations(user_id, event_id) VALUES (?, ?)"
	_, exception := database.Exec(query, userId, event.Id)

	return exception
}

func (event *Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE user_id = ? AND event_id = ?"
	_, exception := database.Exec(query, userId, event.Id)

	return exception
}
