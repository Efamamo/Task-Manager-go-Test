package repository

import (
	"context"
	"time"

	domain "github.com/Task-Management-go/Domain"
	infrastructure "github.com/Task-Management-go/Infrastructure"
	err "github.com/Task-Management-go/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func (ur *UserRepository) Count() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := userCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}

	var user domain.User
	e := userCollection.FindOne(ctx, filter).Decode(&user)

	if e != nil {
		if e == mongo.ErrNoDocuments {
			return nil, err.NewNotFound("User Not Found")
		}
		return nil, err.NewUnexpected(e.Error())
	}

	return &user, nil
}

func (ur *UserRepository) SignUp(user domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hashedPassword, e := infrastructure.HasPassword(user.Password)
	if e != nil {
		return nil, err.NewValidation("Password Cant Be Hashed")
	}
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()
	_, e = userCollection.InsertOne(ctx, user)

	if e != nil {
		return nil, err.NewConflict("user with then given username already exists")
	}

	return &user, nil
}

func (ur *UserRepository) Login(user domain.User) (string, error) {
	existingUser, e := ur.GetUserByUsername(user.Username)

	if e != nil {
		return "", e
	}

	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return "", err.NewUnauthorized("Invalid Credentials")
	}

	jwtToken, err := infrastructure.GenerateToken(existingUser.Username, existingUser.IsAdmin)

	if err != nil {
		return "", e
	}

	return jwtToken, nil
}

func (ur *UserRepository) PromoteUser(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	existingUser, e := ur.GetUserByUsername(username)

	if e != nil {
		return false, err.NewNotFound("User Not Found")
	}

	if existingUser.IsAdmin {
		return false, err.NewConflict("User is already an Admin")
	}

	filter := bson.M{"username": existingUser.Username}
	filter2 := bson.M{"$set": bson.M{"isAdmin": true}}

	_, e = userCollection.UpdateOne(ctx, filter, filter2)

	if e != nil {
		return false, err.NewUnexpected(e.Error())
	}

	return true, nil
}
