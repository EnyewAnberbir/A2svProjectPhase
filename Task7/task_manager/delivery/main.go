package main

import (
	"log"
	"task_manager/delivery/routers"
	"task_manager/infrastructure"
	"task_manager/repositories"
	"task_manager/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("Mongo URI"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(nil)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("task_manager_db")

	taskRepo := repositories.NewTaskRepository(db.Collection("tasks"))
	userRepo := repositories.NewUserRepository(db.Collection("users"), infrastructure.NewPasswordService())
	jwtService := infrastructure.NewJWTService("mySecretKey", "task_manager")
	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	userUseCase := usecases.NewUserUseCase(userRepo)


	r := gin.Default()
	routers.SetupRouter(r, taskUseCase, userUseCase, jwtService)
	r.Run()
}
