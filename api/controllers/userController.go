package controllers

import (
	"context"
	"crud/models"
	"crud/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
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

var validate = validator.New()

func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the user input
	if err := validate.Struct(user); err != nil {
		// If validation fails, return the validation error
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("Field '%s' %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, utils.BadRequest{
			Status:  http.StatusBadRequest,
			Message: "Validation error",
			Errors:  validationErrors,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// If user already exists
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

	// Creating User Picture
	profilePicture := fmt.Sprintf("https://avatar.iran.liara.run/public/boy?username=%s", user.Name)

	user.Password = string(hashPassword)
	user.ProfilePicture = profilePicture

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	user.ID = result.InsertedID.(primitive.ObjectID).Hex()

	type SuccessResponse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    models.User
	}

	return c.JSON(http.StatusOK, SuccessResponse{Status: http.StatusOK, Message: "User created successfully", Data: user})
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
		return c.JSON(http.StatusNotFound, utils.BadRequest{Status: http.StatusNotFound, Message: "Invalid email "})
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

	type SuccessResponse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    models.User
	}
	return c.JSON(http.StatusOK, SuccessResponse{Status: http.StatusOK, Message: "User logged in successfully", Data: existingUser})

}

// Update User
func UpdateUser(c echo.Context) error {
	userID, ok := c.Get("id").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateData := bson.M{}

	if user.Name != "" {
		updateData["name"] = user.Name
	}

	if user.Email != "" {
		updateData["email"] = user.Email
	}

	if user.Password != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
		}
		updateData["password"] = string(hashPassword)
	}

	update := bson.M{"$set": updateData}
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	type SuccessResponse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    models.User
	}

	return c.JSON(http.StatusOK, SuccessResponse{Status: http.StatusOK, Message: "User updated successfully"})
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
		Message: "Users retrieved successfully",
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

	// Remove password before returning response
	userMap := map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "User retrieved successfully",
		Data:    userMap,
	})
}
