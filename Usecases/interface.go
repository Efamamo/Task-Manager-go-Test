package usecases

import (
	domain "github.com/Task-Management-go/Domain"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITaskRepository interface {
	FindAll() (*[]domain.Task, error)
	FindOne(primitive.ObjectID) (*domain.Task, error)
	UpdateOne(primitive.ObjectID, domain.Task) error
	DeleteOne(primitive.ObjectID) (*domain.Task, error)
	Save(domain.Task) (*domain.Task, error)
}

type IUserRepository interface {
	SignUp(user domain.User) (*domain.User, error)
	PromoteUser(username string) (bool, error)
	Count() (int64, error)
	GetUserByUsername(username string) (*domain.User, error)
}

type IPasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(euPassword string, uPassword string) (bool, error)
}

type IJWTService interface {
	ValidateToken(t string) (*jwt.Token, error)
	ValidateAdmin(token *jwt.Token) bool
	GenerateToken(username string, isAdmin bool) (string, error)
}

type IUserService interface {
	SignUp(user domain.User) (*domain.User, error)
	Login(user domain.User) (string, error)
	Promote(username string) (bool, error)
}

type ITaskService interface {
	GetTasks() (*[]domain.Task, error)
	GetTaskByID(id string) (*domain.Task, error)
	UpdateItem(id string, task domain.Task) error
	DeleteTask(id string) (*domain.Task, error)
	AddTask(task domain.Task) (*domain.Task, error)
}
