package controllers

import (
	"context"
	"crud/models"
	"crud/utils"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var clientCollection *mongo.Collection
var userCollect *mongo.Collection

func InitClient(db *mongo.Database) {
	clientCollection = db.Collection("clients")
	userCollect = db.Collection("users")

	if clientCollection == nil {
		log.Println("❌ clientCollection not initialized!")
	} else {
		log.Println("✅ clientCollection initialized successfully!")
	}
}

// Create client
func CreateClient(c echo.Context) error {
	userID, ok := c.Get("id").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var client models.Client
	if err := c.Bind(&client); err != nil {
		return c.JSON(http.StatusBadRequest, utils.BadRequest{Status: http.StatusBadRequest, Message: err.Error()})
	}

	client.CreatedBy, _ = primitive.ObjectIDFromHex(userID)

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

	client.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "Client created successfully",
		Data:    client,
	})
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
func ListClient(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// First convert string createdBy to ObjectID
	pipeline := []bson.M{
		{
			"$addFields": bson.M{
				"createdByObjId": bson.M{
					"$toObjectId": "$createdBy",
				},
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "createdByObjId",
				"foreignField": "_id",
				"as":           "createdByUser",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$createdByUser",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$addFields": bson.M{
				"createdBy": bson.M{
					"name":           "$createdByUser.name",
					"email":          "$createdByUser.email",
					"profilePicture": "$createdByUser.profilePicture",
				},
			},
		},
		{
			"$project": bson.M{
				"createdByObjId": 0,
				"createdByUser":  0,
			},
		},
	}

	cursor, err := clientCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	defer cursor.Close(ctx)

	var clients []bson.M
	if err = cursor.All(ctx, &clients); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "Data retrieved successfully",
		Data:    clients,
	})
}

// Get client by id
func GetClientById(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.BadRequest{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"_id", objID}}}},
		{{"$lookup", bson.D{
			{"from", "users"},
			{"localField", "createdBy"},
			{"foreignField", "_id"},
			{"as", "createdByUser"},
		}}},
		{{"$unwind", bson.D{
			{"path", "$createdByUser"},
			{"preserveNullAndEmptyArrays", true},
		}}},
	}

	cursor, err := clientCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BadRequest{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	defer cursor.Close(ctx)

	var client models.Client
	if cursor.Next(ctx) {
		if err := cursor.Decode(&client); err != nil {
			return c.JSON(http.StatusInternalServerError, utils.BadRequest{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	} else {
		return c.JSON(http.StatusNotFound, utils.BadRequest{
			Status:  http.StatusNotFound,
			Message: "Client not found",
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "Data retrieved successfully",
		Data:    client,
	})
}
