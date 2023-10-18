package routes

import (
	"pos-app/controller"
	"github.com/labstack/echo"
	
)

func Init() *echo.Echo {

	e := echo.New()

	//e = e.Group("/book")
	e.POST("/book", controller.CreateBook)
	e.GET("/book:id", controller.GetBook)
	e.PUT("/book:id", controller.UpdateBook)
	e.DELETE("/book:id", controller.DeleteBook)
	
	return	e
}