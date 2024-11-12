package routes

import (
	"inventory-management-api/src/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterProduct(e *echo.Echo) {
	e.GET("/Inventory/View-Product", handlers.GetProducts)
}
