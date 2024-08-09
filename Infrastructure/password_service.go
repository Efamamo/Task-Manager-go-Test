package infrastructure

import "golang.org/x/crypto/bcrypt"

func HasPassword(password string) (string, error) {
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if e != nil {
		return "", e
	}
	return string(hashedPassword), nil

}
