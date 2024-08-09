package usecases

import (
	domain "github.com/Task-Management-go/Domain"
	repository "github.com/Task-Management-go/Repository"
	err "github.com/Task-Management-go/errors"
)



var userRepository UserInterface = &repository.UserRepository{}

func SignUp(user domain.User) (*domain.User, error) {
	count, e := userRepository.Count()
	if e != nil {
		return nil, err.NewUnexpected(e.Error())
	}

	if count == 0 {
		user.IsAdmin = true
	}
	u, e := userRepository.SignUp(user)

	if e != nil {
		return nil, e
	}

	return u, nil

}

func Login(user domain.User) (string, error) {
	jwtToken, err := userRepository.Login(user)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func Promote(username string) (bool, error) {
	_, err := userRepository.PromoteUser(username)
	if err != nil {
		return false, err
	}
	return true, nil
}
