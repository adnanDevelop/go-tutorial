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
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func Init(db *mongo.Database) {
	userCollection = db.Collection("users")
}

// Create User
func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// If user already exist
	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return c.JSON(http.StatusConflict, utils.BadRequest{Status: http.StatusConflict, Message: "User already exists"})
	}

	// Hashing password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	user.Password = string(hashPassword)

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	user.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return c.JSON(http.StatusOK, utils.Response{Status: http.StatusOK, Message: "User created successfully", Data: user})
}

// Login User
func LoginUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// If user already exist
	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.BadRequest{Status: http.StatusNotFound, Message: "User not found"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.BadRequest{Status: http.StatusUnauthorized, Message: "Invalid password"})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(existingUser.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	c.Response().Header().Add("Authorization", "Bearer "+token)

	type LoginResponse struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	type SuccessResponse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    LoginResponse
	}
	return c.JSON(http.StatusOK, SuccessResponse{Status: http.StatusOK, Message: "User logged in successfully", Data: LoginResponse{ID: existingUser.ID, Name: existingUser.Name, Email: existingUser.Email}})

}

// Update User
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

// Delete User
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

// Get Users
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

// Get User by ID
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
