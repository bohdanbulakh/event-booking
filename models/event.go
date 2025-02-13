package models

import (
	"event-booking/database"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required" json:"name"`
	Description string    `json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"datetime"`
	UserId      int       `binding:"required" json:"user_id"`
}

func (event Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	statement, exception := database.DB.Prepare(query)
	if exception != nil {
		return exception
	}
	defer statement.Close()

	result, exception := statement.Exec(
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

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM events`

	rows, exception := database.DB.Query(query)
	if exception != nil {
		return nil, exception
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		exception = rows.Scan(
			&event.Id,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId,
		)

		if exception != nil {
			return nil, exception
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
	SELECT * FROM events WHERE id = ?`

	row := database.DB.QueryRow(query, id)

	var event Event
	exception := row.Scan(
		&event.Id,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId,
	)

	if exception != nil {
		return nil, exception
	}

	return &event, exception
}

func (event Event) Update() error {
	query := `
UPDATE events
SET name = ?, description = ?, location = ?, datetime = ?, user_id = ?
WHERE id = ?`

	statement, exception := database.DB.Prepare(query)
	if exception != nil {
		return exception
	}
	defer statement.Close()

	_, exception = statement.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserId,
		event.Id,
	)

	return exception
}
