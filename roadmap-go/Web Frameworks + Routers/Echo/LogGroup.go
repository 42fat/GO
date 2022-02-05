package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "secret page")
}

func main() {
	e := echo.New()

	g := e.Group("/admin")
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
	}))
	g.GET("/main", mainAdmin)
	e.Start(":1323")
}
