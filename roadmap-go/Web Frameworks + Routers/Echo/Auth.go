package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func check(c echo.Context) error {
	return c.String(http.StatusOK, "suc")
}

func main() {
	e := echo.New()

	g := e.Group("/group")

	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if username == "jack" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/check", check)
	e.Start(":1323")
}
