package data

import (
	"context"
	"errors"
	"strings"
	"time"

	err "github.com/Task-Management-go/errors"
	model "github.com/Task-Management-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetTasks() (*[]model.Task, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := taskCollection.Find(ctx, bson.D{})
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

	objectId, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		return nil, err.NewValidation("invalid ID format")
	}

	filter := bson.M{"_id": objectId}

	var task model.Task
	e = taskCollection.FindOne(ctx, filter).Decode(&task)
	if e != nil {
		if e == mongo.ErrNoDocuments {
			return nil, err.NewNotFound("task not found")
		}
		return nil, err.NewUnexpected(e.Error())
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
	_, err = taskCollection.UpdateOne(ctx, filter, update)
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
	_, err = taskCollection.DeleteOne(ctx, filter)

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

		_, err := taskCollection.InsertOne(ctx, task)

		if mongo.IsDuplicateKeyError(err) {
			continue
		} else if err != nil {
			return nil, err
		}

		return &task, nil
	}
}
