package main

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"time"
)

type JwtClaims struct {
	Name string `json"name"`
	jwt.StandardClaims
}

func login2(c echo.Context) error {
	username := c.QueryParams()["username"][0]
	password := c.QueryParams()["password"][0]

	if username == "jack" && password == "1234" {
		token, err := createJwtToken()
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "error")
		}
		jwtCookie := &http.Cookie{}
		jwtCookie.Name = "JwtCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(jwtCookie)
		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in!",
			"token":   token,
		})
	}

	return c.String(http.StatusNonAuthoritativeInfo, "You username or password were wrong")
}
func createJwtToken() (string, error) {
	claims := JwtClaims{
		"jack",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Minute).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}
	return token, nil

}
func mainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claim := token.Claims.(jwt.MapClaims)
	log.Println(claim)
	return c.String(http.StatusOK, "It is secret")
}
func main() {
	e := echo.New()

	jwtGroup := e.Group("/jwt")
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
		TokenLookup:   "cookie:JwtCookie",
	}))
	jwtGroup.GET("/main", mainJwt)
	e.GET("/login", login2)
	e.Start(":1323")
}
