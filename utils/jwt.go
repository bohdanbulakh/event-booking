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

func VerifyToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		secret := os.Getenv("SECRET")
		if secret == "" {
			return "", errors.New("cannot find secret in env variables")
		}

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return 0, errors.New("cannot parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	id, ok := claims["id"].(int64)
	if !ok {
		return 0, errors.New("invalid user id")
	}

	return id, nil
}
