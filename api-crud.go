package main

import (
	"api-crud/apicrud"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/products", apicrud.GetProducts)
	e.GET("/product/:id", apicrud.GetProduct)
	e.POST("/product", apicrud.PostProduct)
	e.DELETE("/product/:id", apicrud.DeleteProduct)
	e.PUT("/product/:id", apicrud.PutProduct)

	e.Logger.Fatal(e.Start(":1323"))
}
