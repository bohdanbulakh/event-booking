package models

import "event-booking/database"

type User struct {
	Id       int64
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (user User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)`

	statement, exception := database.DB.Prepare(query)
	if exception != nil {
		return exception
	}

	result, exception := statement.Exec(
		user.Email,
		user.Password,
	)
	if exception != nil {
		return exception
	}

	id, exception := result.LastInsertId()
	user.Id = id
	return exception
}
