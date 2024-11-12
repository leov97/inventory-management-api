package handlers

import (
	"inventory-management-api/src/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

var Products = []models.Procuct{}

func GetProducts(c echo.Context) error {

	return c.JSON(http.StatusOK, Products)

}
