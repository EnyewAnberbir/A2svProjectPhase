package repositories

import (
	"context"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	CreateTask(task domain.Task) (domain.Task, error)
	UpdateTask(id string, task domain.Task) (domain.Task, error)
	DeleteTask(id string) error
	GetTasks() ([]domain.Task, error)
	GetTaskByID(id string) (domain.Task, error)
}

type taskRepository struct {
	db *mongo.Collection
}

func NewTaskRepository(db *mongo.Collection) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) CreateTask(task domain.Task) (domain.Task, error) {
	task.ID = primitive.NewObjectID()
	task.Status = "Pending"
	_, err := r.db.InsertOne(context.Background(), task)
	return task, err
}

func (r *taskRepository) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": task}
	_, err := r.db.UpdateOne(context.Background(), filter, update)
	return task, err
}

func (r *taskRepository) DeleteTask(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	_, err := r.db.DeleteOne(context.Background(), bson.M{"_id": oid})
	return err
}

func (r *taskRepository) GetTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	cursor, err := r.db.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *taskRepository) GetTaskByID(id string) (domain.Task, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	var task domain.Task
	err := r.db.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&task)
	return task, err
}
