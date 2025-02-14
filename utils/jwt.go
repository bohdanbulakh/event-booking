package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(id int64, email string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"email": email,
			"exp": time.
				Now().
				Add(time.Hour * 24).
				Unix(),
		})

	secret := os.Getenv("SECRET")
	if secret == "" {
		return secret, errors.New("cannot find secret in env variables")
	}

	return token.SignedString([]byte(secret))
}
