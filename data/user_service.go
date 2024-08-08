package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	model "github.com/Task-Management-go/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("mini123")

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

func getUserByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}

	var user model.User
	err := userCollection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func Login(user model.User) (string, error) {
	existingUser, err := getUserByUsername(user.Username)

	if err != nil {
		return "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return "", errors.New("invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": existingUser.Username,
	})

	jwtToken, err := token.SignedString(jwtSecret)

	if err != nil {
		fmt.Println(err)
		return "", errors.New("server error")
	}
	return jwtToken, nil
}
