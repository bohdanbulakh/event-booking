package models

import (
	"time"
)

type Event struct {
	Id          int
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"datetime"`
	UserId      int       `binding:"required" json:"user_id"`
}

var events []Event

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
