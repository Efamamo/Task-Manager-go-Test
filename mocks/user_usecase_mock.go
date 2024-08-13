// MockUserService.go
package mocks

import (
	domain "github.com/Task-Management-go/Domain"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) SignUp(user domain.User) (*domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserService) Login(user domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockUserService) Promote(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}
