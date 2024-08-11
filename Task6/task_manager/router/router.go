package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public Routes
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// Protected Routes
	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthenticateJWT())
	{
		authRoutes.GET("/tasks", controllers.GetTasks)
		authRoutes.GET("/tasks/:id", controllers.GetTaskByID)
		authRoutes.POST("/tasks", middleware.AuthorizeAdmin(), controllers.CreateTask)
		authRoutes.PUT("/tasks/:id", middleware.AuthorizeAdmin(), controllers.UpdateTask)
		authRoutes.DELETE("/tasks/:id", middleware.AuthorizeAdmin(), controllers.DeleteTask)
		authRoutes.PUT("/promote/:id", middleware.AuthorizeAdmin(), controllers.PromoteUser)
	}

	return r
}
