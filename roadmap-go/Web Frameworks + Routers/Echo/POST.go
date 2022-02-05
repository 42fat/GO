package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
)

type Cat struct {
	Name string `json"name"`
	Type string `json"type"`
}

type Dog struct {
	Name string `json"name"`
	Type string `json"type"`
}
type Hamsters struct {
	Name string `json"name"`
	Type string `json"type"`
}

func addCat(c echo.Context) error {
	cat := Cat{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error1")
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Println("Error2")
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("this your cat : %#v", cat)
	return c.String(http.StatusOK, "we got cat")
}
func addDog(c echo.Context) error {
	dog := Dog{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed processing addDog request :%s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("this is your dog: %#v", dog)
	return c.String(http.StatusOK, "we got dog")

}
func addHamsters(c echo.Context) error {
	hamster := Hamsters{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed proccessing addHamster request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("this is your dog: %#v", hamster)
	return c.String(http.StatusOK, "we got hamster")
}

func main() {
	e := echo.New()
	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/hamsters", addHamsters)
	e.Start(":1323")
}
