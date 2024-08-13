package repositories

import (
	"context"
	"errors"
	"task_manager/domain"
	"task_manager/infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	AuthenticateUser(username, password string) (domain.User, error)
	PromoteUser(id string) (domain.User, error)
}

type userRepository struct {
	db             *mongo.Collection
	passwordHasher infrastructure.PasswordService
}

func NewUserRepository(db *mongo.Collection, hasher infrastructure.PasswordService) UserRepository {
	return &userRepository{db: db, passwordHasher: hasher}
}

func (r *userRepository) CreateUser(user domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID()
	hashedPassword, err := r.passwordHasher.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = hashedPassword
	user.Role = "user" // Default role
	_, err = r.db.InsertOne(context.Background(), user)
	return user, err
}

func (r *userRepository) AuthenticateUser(username, password string) (domain.User, error) {
	var user domain.User
	err := r.db.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	err = r.passwordHasher.ComparePasswords(user.Password, password)
	if err != nil {
		return domain.User{}, errors.New("invalid credentials")
	}
	return user, nil
}

func (r *userRepository) PromoteUser(id string) (domain.User, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{"role": "admin"}}
	var user domain.User
	err := r.db.FindOneAndUpdate(context.Background(), filter, update).Decode(&user)
	return user, err
}
