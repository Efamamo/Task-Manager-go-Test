package infrastructure

import (
	"github.com/Task-Management-go/Domain/err"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if e != nil {
		return "", e
	}
	return string(hashedPassword), nil
}

func ComparePassword(euPassword string, uPassword string) (bool, error) {

	if bcrypt.CompareHashAndPassword([]byte(euPassword), []byte(uPassword)) != nil {
		return false, err.NewUnauthorized("Invalid Credentials")
	}

	return true, nil
}
