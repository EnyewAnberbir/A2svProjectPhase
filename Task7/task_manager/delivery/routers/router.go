package routers

import (
	"task_manager/delivery/controllers"
	"task_manager/infrastructure"
	"task_manager/usecases"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, taskUseCase usecases.TaskUseCase, userUseCase usecases.UserUseCase, jwtService infrastructure.JWTService) {
	taskController := controllers.NewTaskController(taskUseCase)
	userController := controllers.NewUserController(userUseCase)
	authMiddleware := infrastructure.NewJWTAuthMiddleware(jwtService)

	api := r.Group("/api")
	{
		api.POST("/register", userController.RegisterUser)
		api.POST("/login", userController.LoginUser)

		protected := api.Group("/")
		protected.Use(authMiddleware.AuthenticateJWT())
		{
			protected.POST("/tasks", taskController.CreateTask)
			protected.GET("/tasks", taskController.GetTasks)
			protected.GET("/tasks/:id", taskController.GetTaskByID)
			protected.PUT("/tasks/:id", taskController.UpdateTask)
			protected.DELETE("/tasks/:id", taskController.DeleteTask)

			protected.PUT("/users/:id/promote", authMiddleware.AuthorizeAdmin(), userController.PromoteUser)
		}
	}
}
