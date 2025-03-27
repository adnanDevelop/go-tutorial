package controllers

import (
	"context"
	"crud/models"
	"crud/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

// Init initializes the userCollection with the provided database
func Init(db *mongo.Database) {
	userCollection = db.Collection("users")
}

func GetUsers(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	defer cursor.Close(ctx)

	var users []models.User

	if err = cursor.All(ctx, &users); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    users,
	})
}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.BadRequest{Status: http.StatusBadRequest, Message: err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.BadRequest{Status: http.StatusNotFound, Message: "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	user.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"$set": user}
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, "User updated successfully")
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNotFound, utils.ShortResponse{Status: http.StatusNotFound, Message: "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	_, err = userCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, utils.ShortResponse{Status: http.StatusOK, Message: "User deleted successfully"})
}
