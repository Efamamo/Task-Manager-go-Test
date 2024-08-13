package task_repository_test

import (
	"context"
	"testing"
	"time"

	domain "github.com/Task-Management-go/Domain"
	repository "github.com/Task-Management-go/Repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setup() (*mongo.Client, func(), error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		collection := client.Database("task-management").Collection("tasks")
		collection.DeleteMany(context.TODO(), bson.D{{}})

		client.Disconnect(context.Background())
	}

	return client, cleanup, nil
}

func TestTaskRepository_FindAll(t *testing.T) {
	client, cleanup, err := setup()
	assert.NoError(t, err)
	defer cleanup()

	taskRepo := repository.NewTaskRepo(client)

	// Insert sample data
	task1 := domain.Task{
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "open",
	}
	task2 := domain.Task{
		Title:       "Task 2",
		Description: "Description 2",
		DueDate:     time.Now(),
		Status:      "completed",
	}

	_, err = taskRepo.Save(task1)
	assert.NoError(t, err)
	_, err = taskRepo.Save(task2)
	assert.NoError(t, err)

	tasks, err := taskRepo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, *tasks, 2)
}

func TestTaskRepository_FindOne(t *testing.T) {
	client, cleanup, err := setup()
	assert.NoError(t, err)
	defer cleanup()

	taskRepo := repository.NewTaskRepo(client)

	// Insert a sample task
	task := domain.Task{
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "open",
	}
	savedTask, err := taskRepo.Save(task)
	assert.NoError(t, err)

	retrievedTask, err := taskRepo.FindOne(savedTask.ID)
	assert.NoError(t, err)
	assert.Equal(t, savedTask.ID, retrievedTask.ID)
}

func TestTaskRepository_Save(t *testing.T) {
	client, cleanup, err := setup()
	assert.NoError(t, err)
	defer cleanup()

	taskRepo := repository.NewTaskRepo(client)

	task := domain.Task{
		Title:       "New Task",
		Description: "New Description",
		DueDate:     time.Now(),
		Status:      "open",
	}

	savedTask, err := taskRepo.Save(task)
	assert.NoError(t, err)
	assert.Equal(t, task.Title, savedTask.Title)
}

func TestTaskRepository_UpdateOne(t *testing.T) {
	client, cleanup, err := setup()
	assert.NoError(t, err)
	defer cleanup()

	taskRepo := repository.NewTaskRepo(client)

	task := domain.Task{
		Title:       "Task to Update",
		Description: "Description",
		DueDate:     time.Now(),
		Status:      "open",
	}
	savedTask, err := taskRepo.Save(task)
	assert.NoError(t, err)

	updatedTask := domain.Task{
		Title:       "Updated Title",
		Description: "Updated Description",
		DueDate:     time.Now(),
		Status:      "completed",
	}

	err = taskRepo.UpdateOne(savedTask.ID, updatedTask)
	assert.NoError(t, err)

	retrievedTask, err := taskRepo.FindOne(savedTask.ID)
	assert.NoError(t, err)
	assert.Equal(t, updatedTask.Title, retrievedTask.Title)
	assert.Equal(t, updatedTask.Description, retrievedTask.Description)
	assert.Equal(t, updatedTask.Status, retrievedTask.Status)
}

func TestTaskRepository_DeleteOne(t *testing.T) {
	client, cleanup, err := setup()
	assert.NoError(t, err)
	defer cleanup()

	taskRepo := repository.NewTaskRepo(client)

	task := domain.Task{
		Title:       "Task to Delete",
		Description: "Description",
		DueDate:     time.Now(),
		Status:      "open",
	}
	savedTask, err := taskRepo.Save(task)
	assert.NoError(t, err)

	deletedTask, err := taskRepo.DeleteOne(savedTask.ID)
	assert.NoError(t, err)
	assert.Equal(t, savedTask.ID, deletedTask.ID)

	_, err = taskRepo.FindOne(savedTask.ID)
	assert.Error(t, err)
	assert.Equal(t, "task not found", err.Error())
}
