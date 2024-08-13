// usecases/mocks/mock_task_service.go
package mocks

import (
	domain "github.com/Task-Management-go/Domain"
	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) GetTasks() (*[]domain.Task, error) {
	args := m.Called()
	return args.Get(0).(*[]domain.Task), args.Error(1)
}

func (m *MockTaskService) GetTaskByID(id string) (*domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskService) UpdateItem(id string, task domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskService) DeleteTask(id string) (*domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskService) AddTask(task domain.Task) (*domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*domain.Task), args.Error(1)
}
