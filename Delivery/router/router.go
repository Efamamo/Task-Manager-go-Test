package router

import (
	controller "github.com/Task-Management-go/Delivery/controllers"
	infrastructure "github.com/Task-Management-go/Infrastructure"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(taskController controller.TaskController, userController controller.UserController) {
	r := gin.Default()
	r.GET("/tasks", infrastructure.AuthMiddleware(false), taskController.GetTasks)
	r.GET("tasks/:id", infrastructure.AuthMiddleware(false), taskController.GetTaskById)
	r.PUT("/tasks/:id", infrastructure.AuthMiddleware(true), taskController.UpdateItem)
	r.DELETE("/tasks/:id", infrastructure.AuthMiddleware(true), taskController.DeleteTask)
	r.POST("/tasks", infrastructure.AuthMiddleware(true), taskController.AddTask)
	r.POST("/register", userController.SignUp)
	r.POST("/login", userController.Login)
	r.PATCH("/promote", infrastructure.AuthMiddleware(true), userController.Promote)
	r.Run("localhost:3000")

}
