package repository

import (
	"context"
	"log"
	"time"

	domain "github.com/Task-Management-go/Domain"
	"github.com/Task-Management-go/Domain/err"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection mongo.Collection
}

func NewUserRepo(client *mongo.Client) *UserRepository {
	userCollection := client.Database("task-management").Collection("users")
	cursor, err := userCollection.Indexes().List(context.TODO())
	if err != nil {
		log.Printf("could not list indexes: %v", err)
	}
	defer cursor.Close(context.TODO())

	var indexes []bson.M
	if err := cursor.All(context.TODO(), &indexes); err != nil {
		log.Printf("could not parse indexes: %v", err)
	}

	// Check if the "username" index already exists
	indexExists := false
	for _, index := range indexes {
		key := index["key"].(bson.M)
		if len(key) == 1 && key["username"] != nil {
			indexExists = true
			break
		}
	}

	// If the index does not exist, create it
	if !indexExists {
		indexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}}, // Create index on the "username" field
			Options: options.Index().SetUnique(true),     // Ensure the index is unique
		}

		// Create the index
		_, err = userCollection.Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			log.Printf("could not create index: %v", err)
		}
	} else {
		log.Println("username index already exists")
	}

	return &UserRepository{
		collection: *userCollection,
	}
}

func (ur *UserRepository) Count() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := ur.collection.CountDocuments(ctx, bson.M{})
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
	e := ur.collection.FindOne(ctx, filter).Decode(&user)

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

	user.ID = primitive.NewObjectID()
	_, e := ur.collection.InsertOne(ctx, user)

	if e != nil {
		return nil, err.NewConflict("user with then given username already exists")
	}

	return &user, nil
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

	_, e = ur.collection.UpdateOne(ctx, filter, filter2)

	if e != nil {
		return false, err.NewUnexpected(e.Error())
	}

	return true, nil
}
