package router

import (
	controller "github.com/Task-Management-go/Delivery/controllers"
	infrastructure "github.com/Task-Management-go/Infrastructure"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() {
	r := gin.Default()
	r.GET("/tasks", infrastructure.AuthMiddleware(false), controller.GetTasks)
	r.GET("tasks/:id", infrastructure.AuthMiddleware(false), controller.GetTaskById)
	r.PUT("/tasks/:id", infrastructure.AuthMiddleware(true), controller.UpdateItem)
	r.DELETE("/tasks/:id", infrastructure.AuthMiddleware(true), controller.DeleteTask)
	r.POST("/tasks", infrastructure.AuthMiddleware(true), controller.AddTask)
	r.POST("/register", controller.SignUp)
	r.POST("/login", controller.Login)
	r.PATCH("/promote", infrastructure.AuthMiddleware(true), controller.Promote)
	r.Run("localhost:3000")

}
