package domain

import "go.mongodb.org/mongo-driver/bson/primitive"


type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `json:"username" binding:"required"`
	Password string             `json:"-"`
	Role     string             `json:"role"`
}
