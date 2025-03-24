package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB Collection Variable
var collection *mongo.Collection

// User Struct
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
}

// Pagination struct
type Pagination struct {
	CurrentPage int `json:"currentPage"`
	TotalItems  int `json:"totalItems"`
	TotalPages  int `json:"totalPages"`
}

// Response struct with pagination
type Response struct {
	Status     int        `json:"status"`
	Message    string     `json:"message"`
	Users      []User     `json:"users"`
	Pagination Pagination `json:"pagination"`
}

// Connect to MongoDB
func connectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("testdb").Collection("users")
	fmt.Println("Connected to MongoDB!")
}

// Create User (POST)
func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return c.JSON(http.StatusCreated, user)
}

// Get All Users (GET)
func getUsers(c echo.Context) error {
	// Get query params
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	// Convert to integers (default: page=1, limit=10)
	page, _ := strconv.Atoi(pageParam)
	limit, _ := strconv.Atoi(limitParam)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// Calculate how many documents to skip
	skip := (page - 1) * limit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Count total documents
	totalCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	// Find with pagination
	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetLimit(int64(limit)).SetSkip(int64(skip)))
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	var users []User
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return err
		}
		users = append(users, user)
	}

	// Calculate total pages
	totalPages := (int(totalCount) + limit - 1) / limit

	// Create response with pagination info
	response := &Response{
		Status:  http.StatusOK,
		Message: "Data fetched successfully",
		Users:   users,
		Pagination: Pagination{
			CurrentPage: page,
			TotalItems:  int(totalCount),
			TotalPages:  totalPages,
		},
	}
	fmt.Println(response)
	return c.JSON(http.StatusOK, response)
}

// Update User (PUT)
func updateUser(c echo.Context) error {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return err
	}

	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"name": user.Name, "email": user.Email}}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, bson.M{"message": "User updated successfully"})
}

// Delete User (DELETE)
func deleteUser(c echo.Context) error {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, bson.M{"message": "User deleted successfully"})
}

// Main Function
func main() {
	connectDB()
	e := echo.New()

	// Routes
	e.POST("/users", createUser)
	e.GET("/users", getUsers)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start Server
	e.Start(":8080")
}
