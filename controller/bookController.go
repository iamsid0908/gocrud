package controller

import (
	"net/http"
	"simple-res-api/config"
	"simple-res-api/model"

	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	// b := new(model.Book)
	b := new(model.Book)

	db := config.DB()

	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	book := &model.Book{
		Name:        b.Name,
		Description: b.Description,
	}
	if err := db.Create(&book).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	response := map[string]interface{}{
		"data": b,
	}
	return c.JSON(http.StatusOK, response)

}

// func GetBook(c echo.Context) error {
// 	id := c.Param("id")
// 	db := config.DB()

// 	var books []*model.Book

// 	if res := db.Find(&book, id); res.Error != nil {
// 		data := map[string]interface{}{
// 			"message": res.Error.Error(),
// 		}
// 		return c.JSON(http.StatusOK, data)
// 	}
// 	response := map[string]interface{}{
// 		"data":book[0]
// 	}

// }

func GetAllBooks(c echo.Context) error {
	var books []model.Book
	db := config.DB()

	// Retrieve all books from the database
	if err := db.Find(&books).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	// If no error, return all books
	response := map[string]interface{}{
		"data": books,
	}
	return c.JSON(http.StatusOK, response)
}

// https://medium.com/@fadhlimulyana20/building-rest-api-in-go-with-echo-gorm-and-postgresql-6734cae2b9cf
