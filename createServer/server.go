package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ApiData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		uri := "mongodb://localhost:27017"
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal(err)
		}

		db := client.Database("newData")
		collection := db.Collection("posts")

		var posts []ApiData
		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}

		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var post ApiData
			if err := cursor.Decode(&post); err != nil {
				log.Fatal(err)
			}
			posts = append(posts, post)
		}

		response := echo.Map{
			"status":  http.StatusOK,
			"message": "Get Data successfully",
			"data":    posts,
		}

		return c.JSON(http.StatusOK, response)
	})

	e.POST("/post-data", func(c echo.Context) error {
		var userData PostData

		if err := c.Bind(&userData); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Invalid request format",
			})
		}

		response := echo.Map{
			"status": http.StatusOK,
			"data":   userData,
		}

		return c.JSON(http.StatusOK, response)
	})

	// e.Logger.Fatal(e.Start(":1234"))
	e.Logger.Fatal(e.Start(":1234"))

}
