package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Bluebot/1.0")
		c.Response().Header().Set("MAXIM", "BOBi")
		return next(c)
	}

}

func check1(c echo.Context) error {
	return c.String(http.StatusOK, "OKEY")
}
func main() {
	e := echo.New()
	e.Use(ServerHeader)
	e.GET("", check1)
	e.Start(":1323")
}
