package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, exception := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	return string(bytes), exception
}

func Compare(unHashed string, hashed string) bool {
	exception := bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(unHashed),
	)
	return exception == nil
}
