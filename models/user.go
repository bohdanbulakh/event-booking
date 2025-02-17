package models

import (
	"database/sql"
	"errors"
	"event-booking/database"
	"event-booking/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (user *User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)`

	hashedPassword, exception := utils.HashPassword(user.Password)
	if exception != nil {
		return exception
	}

	result, exception := database.Exec(
		query,
		user.Email,
		hashedPassword,
	)

	if exception != nil {
		return exception
	}

	id, exception := result.LastInsertId()
	user.Id = id
	return exception
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	var hashedPassword string

	_, exception := database.QueryRow(query, func(row *sql.Row) error {
		return row.Scan(
			&user.Id,
			&hashedPassword,
		)
	}, user.Email)

	if exception != nil {
		return exception
	}

	isPasswordValid := utils.Compare(
		user.Password,
		hashedPassword,
	)

	if !isPasswordValid {
		return errors.New("password is incorrect")
	}
	return nil
}
