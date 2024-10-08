package repository_test

import (
	"context"
	"testing"
	"time"

	domain "github.com/Task-Management-go/Domain"
	repository "github.com/Task-Management-go/Repository"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	client   *mongo.Client
	taskRepo *repository.TaskRepository
	cleanup  func()
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	suite.Require().NoError(err)

	suite.client = client
	suite.taskRepo = repository.NewTaskRepo(client)

	suite.cleanup = func() {
		collection := client.Database("task-management").Collection("tasks")
		collection.DeleteMany(context.TODO(), bson.D{{}})
		client.Disconnect(context.Background())
	}
}

func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	suite.cleanup()
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}

func (suite *TaskRepositoryTestSuite) TestTaskRepository_FindAll() {
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

	_, err := suite.taskRepo.Save(task1)
	suite.Require().NoError(err)
	_, err = suite.taskRepo.Save(task2)
	suite.Require().NoError(err)

	tasks, err := suite.taskRepo.FindAll()
	suite.Require().NoError(err)
	suite.Require().Len(*tasks, 2)
}

func (suite *TaskRepositoryTestSuite) TestTaskRepository_FindOne() {
	// Insert a sample task
	task := domain.Task{
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "open",
	}
	savedTask, err := suite.taskRepo.Save(task)
	suite.Require().NoError(err)

	retrievedTask, err := suite.taskRepo.FindOne(savedTask.ID)
	suite.Require().NoError(err)
	suite.Require().Equal(savedTask.ID, retrievedTask.ID)
}

func (suite *TaskRepositoryTestSuite) TestTaskRepository_Save() {
	task := domain.Task{
		Title:       "New Task",
		Description: "New Description",
		DueDate:     time.Now(),
		Status:      "open",
	}

	savedTask, err := suite.taskRepo.Save(task)
	suite.Require().NoError(err)
	suite.Require().Equal(task.Title, savedTask.Title)
}

func (suite *TaskRepositoryTestSuite) TestTaskRepository_UpdateOne() {
	task := domain.Task{
		Title:       "Task to Update",
		Description: "Description",
		DueDate:     time.Now(),
		Status:      "open",
	}
	savedTask, err := suite.taskRepo.Save(task)
	suite.Require().NoError(err)

	updatedTask := domain.Task{
		Title:       "Updated Title",
		Description: "Updated Description",
		DueDate:     time.Now(),
		Status:      "completed",
	}

	err = suite.taskRepo.UpdateOne(savedTask.ID, updatedTask)
	suite.Require().NoError(err)

	retrievedTask, err := suite.taskRepo.FindOne(savedTask.ID)
	suite.Require().NoError(err)
	suite.Require().Equal(updatedTask.Title, retrievedTask.Title)
	suite.Require().Equal(updatedTask.Description, retrievedTask.Description)
	suite.Require().Equal(updatedTask.Status, retrievedTask.Status)
}

func (suite *TaskRepositoryTestSuite) TestTaskRepository_DeleteOne() {
	task := domain.Task{
		Title:       "Task to Delete",
		Description: "Description",
		DueDate:     time.Now(),
		Status:      "open",
	}
	savedTask, err := suite.taskRepo.Save(task)
	suite.Require().NoError(err)

	deletedTask, err := suite.taskRepo.DeleteOne(savedTask.ID)
	suite.Require().NoError(err)
	suite.Require().Equal(savedTask.ID, deletedTask.ID)

	_, err = suite.taskRepo.FindOne(savedTask.ID)
	suite.Require().Error(err)
	suite.Require().Equal("task not found", err.Error())
}
