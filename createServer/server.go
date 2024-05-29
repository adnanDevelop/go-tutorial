package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {

	var myMap = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	fmt.Println(myMap)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
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

	e.Logger.Fatal(e.Start(":1234"))

}
