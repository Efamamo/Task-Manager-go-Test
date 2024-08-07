package controller

import (
	"fmt"
	"net/http"

	"github.com/Task-Management-go/data"
	model "github.com/Task-Management-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetTasks(c *gin.Context) {
	tasks := data.GetTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task := data.GetTaskById(id)

	if task == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, *task)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedTask model.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}

		errorMessages := make(map[string]string)
		for _, e := range validationErrors {
			field := e.Field()
			switch field {
			case "Title":
				errorMessages["title"] = "Title is required."
			case "Description":
				errorMessages["description"] = "Description is required."

			case "Status":
				errorMessages["status"] = "Status is required."
			}

		}

		// Return a 400 Bad Request response with detailed error messages
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	task, err := data.UpdateItem(id, updatedTask)

	if err != nil && err.Error() == "Status Error" {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Status can only be Pending or In Progress or Completed"})
		return
	}

	if err != nil && err.Error() == "Not Found" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, *task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task := data.DeleteTask(id)

	if task == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusAccepted, *task)
}

func AddTask(c *gin.Context) {
	var newTask model.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}

		errorMessages := make(map[string]string)
		for _, e := range validationErrors {

			field := e.Field()
			fmt.Println(field, "this is field")
			switch field {
			case "Title":
				errorMessages["title"] = "Title is required."
			case "Description":
				errorMessages["description"] = "Description is required."

			case "Status":
				errorMessages["status"] = "Status is required."

			case "DueDate":
				errorMessages["due_date"] = "DueDate is required."
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	task, err := data.AddTask(newTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Status can only be Pending or In Progress or Completed"})
		return
	}

	c.IndentedJSON(http.StatusCreated, task)
}
