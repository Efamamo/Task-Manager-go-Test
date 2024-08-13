// controller/task_controller_test.go
package taskcontrollertest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	controller "github.com/Task-Management-go/Delivery/controllers"
	domain "github.com/Task-Management-go/Domain"
	"github.com/Task-Management-go/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerTestSuite struct {
	suite.Suite
	mockTaskService *mocks.MockTaskService
	taskController  *controller.TaskController
	router          *gin.Engine
}

func (suite *TaskControllerTestSuite) SetupTest() {
	suite.mockTaskService = new(mocks.MockTaskService)
	suite.taskController = &controller.TaskController{Service: suite.mockTaskService}

	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
}

func (suite *TaskControllerTestSuite) TestGetTasks_Success() {
	t1, _ := time.Parse(time.RFC3339Nano, "2024-08-08T09:41:00.564241718+03:00")
	t2, _ := time.Parse(time.RFC3339Nano, "2024-08-08T09:41:00.564241718+05:00")
	task1Id, _ := primitive.ObjectIDFromHex("66b90802a6d354d828cd8f1a")
	task2Id, _ := primitive.ObjectIDFromHex("66b918ff59c53bc04a1d2594")
	tasks := []domain.Task{
		{ID: task1Id, Title: "Task 1", Description: "Description 1", Status: "Pending", DueDate: t1},
		{ID: task2Id, Title: "Task 2", Description: "Description 2", Status: "Completed", DueDate: t2},
	}
	suite.mockTaskService.On("GetTasks").Return(&tasks, nil)

	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.GET("/tasks", suite.taskController.GetTasks)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
	suite.mockTaskService.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestGetTaskByID_Success() {
	t, _ := time.Parse(time.RFC3339Nano, "2024-08-08T09:41:00.564241718+03:00")
	taskId, _ := primitive.ObjectIDFromHex("66b90802a6d354d828cd8f1a")
	task := &domain.Task{ID: taskId, Title: "Task 1", Description: "Description 1", Status: "Pending", DueDate: t}
	suite.mockTaskService.On("GetTaskByID", "1").Return(task, nil)

	req, err := http.NewRequest(http.MethodGet, "/tasks/1", nil)
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.GET("/tasks/:id", suite.taskController.GetTaskById)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
	suite.mockTaskService.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestAddTask_Success() {
	taskId, _ := primitive.ObjectIDFromHex("66b90802a6d354d828cd8f1a")
	t, _ := time.Parse(time.RFC3339Nano, "2024-08-08T09:41:00.564241718+03:00")
	newTask := domain.Task{ID: taskId, Title: "New Task", Description: "New Description", Status: "Pending", DueDate: t}
	suite.mockTaskService.On("AddTask", newTask).Return(&newTask, nil)

	body, _ := json.Marshal(newTask)
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.POST("/tasks", suite.taskController.AddTask)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusCreated, rr.Code)
	suite.mockTaskService.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestUpdateItem_Success() {
	taskId, _ := primitive.ObjectIDFromHex("66b90802a6d354d828cd8f1a")
	t, _ := time.Parse(time.RFC3339Nano, "2024-08-08T09:41:00.564241718+03:00")
	updatedTask := domain.Task{ID: taskId, Title: "New Task", Description: "New Description", Status: "Pending", DueDate: t}
	suite.mockTaskService.On("UpdateItem", "66b90802a6d354d828cd8f1a", updatedTask).Return(nil)

	body, _ := json.Marshal(updatedTask)
	req, err := http.NewRequest(http.MethodPut, "/tasks/66b90802a6d354d828cd8f1a", bytes.NewBuffer(body))
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.PUT("/tasks/:id", suite.taskController.UpdateItem)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
	suite.mockTaskService.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestDeleteTask_Success() {
	taskId, _ := primitive.ObjectIDFromHex("66b90802a6d354d828cd8f1a")
	t, _ := time.Parse(time.RFC3339Nano, "2024-08-08T09:41:00.564241718+03:00")
	task := &domain.Task{ID: taskId, Title: "Task 1", Description: "Description 1", Status: "Pending", DueDate: t}
	suite.mockTaskService.On("DeleteTask", "1").Return(task, nil)

	req, err := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	suite.router.DELETE("/tasks/:id", suite.taskController.DeleteTask)
	suite.router.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusAccepted, rr.Code)
	suite.mockTaskService.AssertExpectations(suite.T())
}

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
