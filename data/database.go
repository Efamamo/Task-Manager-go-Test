package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var taskCollection *mongo.Collection
var userCollection *mongo.Collection

// Initialize MongoDB connection once
func init() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:2717")

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	taskCollection = client.Database("task-management").Collection("tasks")
	userCollection = client.Database("task-management").Collection("users")

	
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"username": 1}, // Index on email field
		Options: options.Index().SetUnique(true),
	}

	// Create the index
	indexName, err := userCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created index:", indexName)
}
