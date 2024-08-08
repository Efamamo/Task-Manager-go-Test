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

	collection = client.Database("task-management").Collection("tasks")
	fmt.Println("Connected to MongoDB!")

}

func GetTasks() (*[]model.Task, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	tasks := make([]model.Task, 0)
	for cur.Next(ctx) {
		var task model.Task

		err := cur.Decode(&task)

		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if cur.Err() != nil {
		return nil, err
	}

	cur.Close(ctx)

	return &tasks, nil

}

func GetTaskByID(id string) (*model.Task, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectId}

	var task model.Task
	err = collection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
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

	_, err = GetTaskByID(ID)
	if err != nil {
		return nil, err
	}

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

	for {
		task.ID = primitive.NewObjectID()

		_, err := collection.InsertOne(ctx, task)

		if mongo.IsDuplicateKeyError(err) {
			continue
		} else if err != nil {
			return nil, err
		}

		return &task, nil
	}
}
