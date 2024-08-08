package controller

import (
	"fmt"
	"net/http"

	"github.com/Task-Management-go/data"
	model "github.com/Task-Management-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SignUp(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}

		errorMessages := make(map[string]string)
		for _, e := range validationErrors {

			field := e.Field()
			fmt.Println(field, "this is field")
			switch field {
			case "Username":
				errorMessages["username"] = "username is required."
			case "Password":
				errorMessages["password"] = "Password is required."

			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	u, err := data.SignUp(newUser)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err})
		return
	}
	c.IndentedJSON(201, u)

}

func Login(c *gin.Context) {
	var reqUser model.User
	if err := c.ShouldBindJSON(&reqUser); err != nil {

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}

		errorMessages := make(map[string]string)
		for _, e := range validationErrors {

			field := e.Field()
			fmt.Println(field, "this is field")
			switch field {
			case "Username":
				errorMessages["username"] = "username is required."
			case "Password":
				errorMessages["password"] = "Password is required."

			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	u, err := data.Login(reqUser)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(201, u)

}
