package data

import (
	"context"
	"time"

	err "github.com/Task-Management-go/errors"
	model "github.com/Task-Management-go/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret = []byte("mini123")

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
	count, e := countUsers()

	if e != nil {
		return nil, err.NewUnexpected(e.Error())
	}

	if count == 0 {
		user.IsAdmin = true
	}

	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

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

func getUserByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}

	var user model.User
	e := userCollection.FindOne(ctx, filter).Decode(&user)

	if e != nil {
		if e == mongo.ErrNoDocuments {
			return nil, err.NewNotFound("User Not Found")
		}
		return nil, err.NewUnexpected(e.Error())
	}

	return &user, nil
}

func Login(user model.User) (string, error) {
	existingUser, e := getUserByUsername(user.Username)

	if e != nil {
		return "", err.NewNotFound("User Not Found")
	}

	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return "", err.NewUnauthorized("Invalid Credentials")
	}

	expirationTime := time.Now().Add(1 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": existingUser.Username,
		"role":     existingUser.IsAdmin,
		"exp":      expirationTime,
	})

	jwtToken, e := token.SignedString(JwtSecret)

	if e != nil {

		return "", err.NewUnexpected(e.Error())
	}
	return jwtToken, nil
}

func Promote(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	existingUser, e := getUserByUsername(username)

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
