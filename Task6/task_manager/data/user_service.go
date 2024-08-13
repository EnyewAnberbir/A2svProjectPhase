package data

import (
	"context"
	"errors"
	"log"
	"time"
	"task_manager/models"
	"os"

	"github.com/joho/godotenv"

	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

func ConnectUserDB() {
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
	    }
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

	userCollection = client.Database("task_manager_db").Collection("users")
	log.Println("Connected to MongoDB (Users)!")
}

func RegisterUser(user models.User) (models.User, error) {
	var existingUser models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the username already exists
	err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return models.User{}, errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	// Set default role if not provided
	if user.Role == "" {
		user.Role = "user"
	}

	// Set the hashed password and ID
	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)

	// Insert the new user into the collection
	if _, err := userCollection.InsertOne(ctx, user); err != nil {
		return models.User{}, err
	}

	return user, nil
}



func LoginUser(username, password string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("incorrect password")
	}

	return user, nil
}

func PromoteUser(id primitive.ObjectID) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userCollection.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"role": "admin"}}).Decode(&user)
	if err != nil {
		return user, err
	}

	user.Role = "admin"
	return user, nil
}
