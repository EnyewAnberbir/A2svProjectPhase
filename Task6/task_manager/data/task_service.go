package data

import (
	"context"
	"log"
	"time"
	"task_manager/models"
	"os"
	
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var taskCollection *mongo.Collection

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Get the MongoDB URI from the environment variable
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        log.Fatal("MONGO_URI environment variable not set")
    }
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	taskCollection = client.Database("task_manager_db").Collection("tasks")
	log.Println("Connected to MongoDB!")
}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.M{})
	if err != nil {
		return tasks, err
	}

	if err = cursor.All(ctx, &tasks); err != nil {
		return tasks, err
	}

	return tasks, nil
}

func GetTaskByID(id primitive.ObjectID) (models.Task, error) {
	var task models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := taskCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return task, err
	}

	return task, nil
}

func CreateTask(task models.Task) (models.Task, error) {
	task.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := taskCollection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}

	return task, nil
}

func UpdateTask(id primitive.ObjectID, updatedTask models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := taskCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updatedTask})
	if err != nil {
		return updatedTask, err
	}

	return GetTaskByID(id)
}

func DeleteTask(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := taskCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
