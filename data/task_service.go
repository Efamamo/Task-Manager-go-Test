package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	model "github.com/Task-Management-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

// Initialize MongoDB connection once
func init() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://nest:efamamo@cluster0.avreuwg.mongodb.net/")

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

	collection = client.Database("task-management").Collection("tasks")
	fmt.Println("Connected to MongoDB!")
}

func GetTasks() (*[]model.Task, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {

		return nil, err
	}

	defer cursor.Close(ctx)

	var tasks []model.Task
	for cursor.Next(ctx) {
		var task model.Task
		if err := cursor.Decode(&task); err != nil {

			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {

		return nil, err
	}

	return &tasks, nil
}

func GetTaskByID(id string) (*model.Task, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Convert the id string to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Create a filter to match the specific ID
	filter := bson.M{"_id": objectID}

	// Find a single document that matches the filter
	var task model.Task
	err = collection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task Not Found")
		}
		return nil, err
	}

	return &task, nil

}

func UpdateItem(ID string, updatedTask model.Task) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	task, err := GetTaskByID(ID)
	if err != nil {
		return nil, err
	}
	fmt.Println(task)

	if strings.ToLower(updatedTask.Status) != "in progress" && strings.ToLower(updatedTask.Status) != "completed" && strings.ToLower(updatedTask.Status) != "pending" {
		return nil, errors.New("status error")
	}
	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	// Update the document that matches the filter
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	utask, err := GetTaskByID(ID)

	if err != nil {
		return nil, err
	}

	return utask, nil

}

func DeleteTask(ID string) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	task, err := GetTaskByID(ID)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}

	// Delete the document that matches the filter
	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return task, nil

}

func AddTask(task model.Task) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if strings.ToLower(task.Status) != "in progress" && strings.ToLower(task.Status) != "completed" && strings.ToLower(task.Status) != "pending" {
		return nil, errors.New("status error")
	}

	task.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		log.Fatal(err)
	}

	return &task, nil
}
