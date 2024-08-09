package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	domain "github.com/Task-Management-go/Domain"
	err "github.com/Task-Management-go/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct{}

func (tr *TaskRepository) FindAll() (*[]domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := taskCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	tasks := make([]domain.Task, 0)
	for cur.Next(ctx) {
		var task domain.Task

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

func (tr *TaskRepository) FindOne(id string) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objectId, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		return nil, err.NewValidation("invalid ID format")
	}

	filter := bson.M{"_id": objectId}

	var task domain.Task
	e = taskCollection.FindOne(ctx, filter).Decode(&task)
	if e != nil {
		if e == mongo.ErrNoDocuments {
			return nil, err.NewNotFound("task not found")
		}
		return nil, err.NewUnexpected(e.Error())
	}

	return &task, nil
}

func (tr *TaskRepository) UpdateOne(id string, updatedTask domain.Task) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, e := tr.FindOne(id)
	if e != nil {
		return nil, e
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(updatedTask.Status) != "in progress" && strings.ToLower(updatedTask.Status) != "completed" && strings.ToLower(updatedTask.Status) != "pending" {
		return nil, errors.New("status error")
	}
	filter := bson.M{"_id": ID}

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

	utask, err := tr.FindOne(id)

	if err != nil {
		return nil, err
	}

	return utask, nil

}

func (tr *TaskRepository) DeleteOne(ID string) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	task, err := tr.FindOne(ID)

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

func (tr *TaskRepository) Save(task domain.Task) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	

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
