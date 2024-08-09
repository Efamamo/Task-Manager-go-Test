package usecases

import domain "github.com/Task-Management-go/Domain"

type TaskInterface interface {
	FindAll() (*[]domain.Task, error)
	FindOne(string) (*domain.Task, error)
	UpdateOne(id string, updatedTask domain.Task) (*domain.Task, error)
	DeleteOne(string) (*domain.Task, error)
	Save(domain.Task) (*domain.Task, error)
}

type UserInterface interface {
	SignUp(user domain.User) (*domain.User, error)
	Login(user domain.User) (string, error)
	PromoteUser(username string) (bool, error)
	Count() (int64, error)
}
