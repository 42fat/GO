package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getCats(c echo.Context) error {
	Param := c.QueryParams()

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("You cat name is: %s\nand his type is: %s\n", Param["name"][0], Param["type"][0]))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": Param["name"][0],
			"type": Param["type"][0],
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "bad request",
	})

}

func main() {
	e := echo.New()
	e.GET("/cats/:data", getCats)
	e.Start(":1323")
}
