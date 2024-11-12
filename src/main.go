package main

import (
	"inventory-management-api/src/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.RegisterProduct(e)

	log.Fatal(e.Start(":8080"))
}
