package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
	"time"
)

func mainCoockie(c echo.Context) error {
	return c.String(http.StatusOK, "COOKIES")
}
func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you dont have cookie")
			}
			log.Println(err)
			return err
		}
		if cookie.Value == "some_string" {
			return next(c)
		}
		return err
	}
}

func login(c echo.Context) error {
	username := c.QueryParams()["username"][0]
	password := c.QueryParams()["password"][0]

	if username == "jack" && password == "1234" {
		cookie := &http.Cookie{}

		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(60 * time.Minute)
		c.SetCookie(cookie)
		return c.String(http.StatusOK, "You were logged in!")
	}
	return c.String(http.StatusNonAuthoritativeInfo, "You username or password were wrong")
}
func main() {
	e := echo.New()
	cookies := e.Group("/cookie")
	cookies.Use(checkCookie)
	cookies.GET("/main", mainCoockie)

	e.GET("/login", login)
	e.Start(":1323")
}
