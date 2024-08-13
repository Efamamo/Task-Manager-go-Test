package controller

import (
	"fmt"
	"net/http"

	domain "github.com/Task-Management-go/Domain"
	"github.com/Task-Management-go/Domain/err"
	usecases "github.com/Task-Management-go/Usecases"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	Service usecases.IUserService
}

func (uc *UserController) SignUp(c *gin.Context) {

	var newUser domain.User
	if err := c.ShouldBindJSON(&newUser); err != nil {

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
			validationErrors = errors
		}

		errorMessages := make(map[string]string)
		for _, e := range validationErrors {

			field := e.Field()

			switch field {
			case "Username":
				errorMessages["username"] = "Username is required."
			case "Password":
				errorMessages["password"] = "Password is required."

			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	u, e := uc.Service.SignUp(newUser)
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

	c.IndentedJSON(201, u)

}

func (uc *UserController) Login(c *gin.Context) {
	var reqUser domain.User
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

	u, e := uc.Service.Login(reqUser)
	if e != nil {
		if e.(*err.Error).Type() == "ServerError" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "NotFound" {
			c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Conflict" {
			c.JSON(http.StatusConflict, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Validation" {
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		} else if e.(*err.Error).Type() == "Unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": e.Error()})
		}
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", u, 1200, "", "", false, true)
	c.IndentedJSON(201, gin.H{"token": u})

}

func (uc *UserController) Promote(c *gin.Context) {
	username := c.Query("username")

	promoted, e := uc.Service.Promote(username)

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

	if promoted {
		c.IndentedJSON(200, gin.H{})
	}

}
