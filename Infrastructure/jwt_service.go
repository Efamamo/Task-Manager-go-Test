package infrastructure

import (
	"fmt"
	"os"
	"strings"
	"time"

	err "github.com/Task-Management-go/errors"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(authHeader string) (*jwt.Token, error) {

	if authHeader == "" {
		return nil, err.NewUnauthorized("unauthorized")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return nil, err.NewUnauthorized("unauthorized")
	}

	token, e := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
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
	fmt.Println(os.Getenv("JwtSecret"))
	jwtToken, e := token.SignedString([]byte(os.Getenv("JwtSecret")))

	if e != nil {
		return "", err.NewValidation("Cant Sign Token")
	}

	return jwtToken, nil
}
