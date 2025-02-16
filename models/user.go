package models

import (
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

	statement, exception := database.DB.Prepare(query)
	if exception != nil {
		return exception
	}

	hashedPassword, exception := utils.HashPassword(user.Password)
	if exception != nil {
		return exception
	}

	result, exception := statement.Exec(
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
	row := database.DB.QueryRow(query, user.Email)

	var hashedPassword string
	exception := row.Scan(
		&user.Id,
		&hashedPassword,
	)
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
