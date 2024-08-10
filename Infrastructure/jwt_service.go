package infrastructure

import (
	"fmt"
	"os"
	"time"

	"github.com/Task-Management-go/Domain/err"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(t string) (*jwt.Token, error) {
	token, e := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JwtSecret")), nil
	})

	if e != nil || !token.Valid {
		return nil, err.NewUnauthorized("unauthorized")
	}
	return token, nil
}

func ValidateAdmin(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}

	role, ok := claims["isAdmin"].(bool)
	if !ok || !role {

		return false
	}
	return true
}

func GenerateToken(username string, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"isAdmin":  isAdmin,
		"exp":      expirationTime,
	})

	jwtToken, e := token.SignedString([]byte(os.Getenv("JwtSecret")))

	if e != nil {
		return "", err.NewValidation("Cant Sign Token")
	}

	return jwtToken, nil
}
