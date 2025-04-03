package controllers

import (
	"context"
	"crud/models"
	"crud/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var clientCollection *mongo.Collection

func InitClient(db *mongo.Database) {
	clientCollection = db.Collection("clients")
}

// Create client
func CreateClient(c echo.Context) error {
	var client models.Client
	fmt.Println(client, "client data")
	if err := c.Bind(&client); err != nil {
		return c.JSON(http.StatusBadRequest, utils.BadRequest{Status: http.StatusBadRequest, Message: err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingClient models.Client
	err := clientCollection.FindOne(ctx, bson.M{"email": client.Email}).Decode(&existingClient)

	if err == nil {
		return c.JSON(http.StatusConflict, utils.BadRequest{Status: http.StatusConflict, Message: "Client already exists"})
	}

	result, err := clientCollection.InsertOne(ctx, client)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	client.ID, _ = result.InsertedID.(primitive.ObjectID)
	return c.JSON(http.StatusOK, utils.Response{Status: http.StatusOK, Message: "Client created successfully", Data: client})

}

// Update Client

// Delete Client
func DeleteClient(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingClient models.Client
	err = clientCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&existingClient)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNotFound, utils.ShortResponse{Status: http.StatusNotFound, Message: "Client not found"})
		}
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	_, err = clientCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, utils.ShortResponse{Status: http.StatusOK, Message: "Client deleted successfully"})

}

// List Client

// Get client by id
