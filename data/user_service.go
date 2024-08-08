package data

import (
	"context"
	"fmt"
	"time"

	model "github.com/Task-Management-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func countUsers() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := userCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}

func SignUp(user model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, err := countUsers()

	if err != nil {
		return nil, err
	}

	if count == 0 {
		user.IsAdmin = true
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	user.ID = primitive.NewObjectID()
	fmt.Println(user)
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
