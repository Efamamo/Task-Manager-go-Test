package controller

import (
	"fmt"
	"net/http"

	domain "github.com/Task-Management-go/Domain"
	usecases "github.com/Task-Management-go/Usecases"

	err "github.com/Task-Management-go/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetTasks(c *gin.Context) {

	tasks, e := usecases.GetTasks()

	if e != nil {
		if e.(*err.Error).Type() == "ServerError" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "NotFound" {
			c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Conflict" {
			c.JSON(http.StatusConflict, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Validation" {
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, e := usecases.GetTaskByID(id)

	if e != nil {
		if e.(*err.Error).Type() == "ServerError" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "NotFound" {
			c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Conflict" {
			c.JSON(http.StatusConflict, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Validation" {
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, *task)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedTask domain.Task
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

	task, e := usecases.UpdateItem(id, updatedTask)

	if e != nil {
		if e.Error() == "status error" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Status can only be Pending or In Progress or Completed"})
		} else if e.(*err.Error).Type() == "ServerError" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "NotFound" {
			c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Conflict" {
			c.JSON(http.StatusConflict, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Validation" {
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, *task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task, e := usecases.DeleteTask(id)

	if e != nil {
		if e.(*err.Error).Type() == "ServerError" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "NotFound" {
			c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Conflict" {
			c.JSON(http.StatusConflict, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Validation" {
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusAccepted, *task)
}

func AddTask(c *gin.Context) {
	var newTask domain.Task

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
		if len(errorMessages) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "DueDate is required."})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	task, e := usecases.AddTask(newTask)

	if e != nil {
		if e.Error() == "status error" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Status can only be Pending or In Progress or Completed"})
		} else if e.(*err.Error).Type() == "ServerError" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "NotFound" {
			c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Conflict" {
			c.JSON(http.StatusConflict, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Validation" {
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusCreated, task)
}
