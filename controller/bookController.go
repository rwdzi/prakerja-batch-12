package controller

import (
	"net/http"
	"pos-app/config"
	"pos-app/models"

	"github.com/labstack/echo"
)

func CreateBook(c echo.Context) error {
	b := new(models.Book)
	db := config.DB()

	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		
		return c.JSON(http.StatusInternalServerError, data)

	}

	book := &models.Book{
		Name:		b.Name,
		Address:	b.Address,
		Phone:		b.Phone,
	}

	if err := db.Create(&book).Error; err != nil {
		data := map[string]interface{} {
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{} {
		"data": b,
	}

	return c.JSON(http.StatusOK, response)

}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	b := new(models.Book)
	db := config.DB()

	if err := c.Bind(b); err != nil {
		data := map[string]interface{} {
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)

	}

	existing_book := new(models.Book)

	if err := db.First(&existing_book, id).Error; err != nil {
		data:= map[string]interface{} {
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)

	}

	existing_book.Name = b.Name
	existing_book.Address = b.Address
	existing_book.Phone = b.Phone

	if err := db.Save(&existing_book).Error; err != nil {
		data := map[string]interface{} {
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)

	}

	response := map[string]interface{} {
		"data": existing_book,
	}

	return c.JSON(http.StatusOK, response)

}

func GetBook(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var book []*models.Book

	if res := db.Find(&book, id); res.Error != nil {
		data := map[string]interface{} {
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{} {
		"data": book[0],
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	book := new(models.Book)

	err := db.Delete(&book, id).Error

	if err != nil {
		data := map[string]interface{} {
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{} {
		"message": "a book has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}