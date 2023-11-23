package main

import (
	"net/http"
	"simple-res-api/config"
	"simple-res-api/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// b := new(model.Book)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	bookRoute := e.Group("/book")
	bookRoute.GET("/", controller.GetAllBooks)
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()
	e.Logger.Fatal(e.Start(":8080"))
}
