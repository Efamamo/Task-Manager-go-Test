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

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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

	if err := ensureIndex(client, "task-management", "users", "username"); err != nil {
		log.Fatal(err)
	}
}

func ensureIndex(client *mongo.Client, databaseName, collectionName, fieldName string) error {
	userCollection := client.Database(databaseName).Collection(collectionName)

	// List the existing indexes
	cursor, err := userCollection.Indexes().List(context.TODO())
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	// Check if the index already exists
	indexExists := false
	for cursor.Next(context.TODO()) {
		var index bson.M
		if err := cursor.Decode(&index); err != nil {
			return err
		}
		if key, ok := index["key"].(bson.M); ok {
			if _, exists := key[fieldName]; exists {
				indexExists = true
				break
			}
		}
	}

	if indexExists {
		fmt.Println("Index on", fieldName, "already exists.")
		return nil
	}

	// Create the index if it does not exist
	indexModel := mongo.IndexModel{
		Keys:    bson.M{fieldName: 1}, // Index on the specified field
		Options: options.Index().SetUnique(true),
	}

	indexName, err := userCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}

	fmt.Println("Created index:", indexName)
	return nil
}
