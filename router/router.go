package router

import (
	controller "github.com/Task-Management-go/controllers"
	"github.com/Task-Management-go/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() {
	r := gin.Default()
	r.GET("/tasks", middleware.AuthMiddleware(false), controller.GetTasks)
	r.GET("tasks/:id", middleware.AuthMiddleware(false), controller.GetTaskById)
	r.PUT("/tasks/:id", middleware.AuthMiddleware(true), controller.UpdateItem)
	r.DELETE("/tasks/:id", middleware.AuthMiddleware(true), controller.DeleteTask)
	r.POST("/tasks", middleware.AuthMiddleware(true), controller.AddTask)
	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)
	r.PATCH("/promote", middleware.AuthMiddleware(true), controller.Promote)
	r.Run("localhost:3000")

}
