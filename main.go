package main

import (
	"marketplace/cmd/handlers"
	"marketplace/front"

	//"marketplace/cmd/repositories"
	"marketplace/cmd/storage"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/", "assets")

	storage.InitDB()
	apiGroup := e.Group("/api")
	//apiGroup.POST("/users", handlers.CreateUser)

	apiGroup.POST("/products", handlers.CreateProduct)

	apiGroup.PUT("/products/:id", handlers.HandleUpdateProduct)

	apiGroup.GET("/products", handlers.HandleGetAllProducts)
	apiGroup.GET("/products/:id", handlers.HandleGetProduct)

	apiGroup.DELETE("/products/:id", handlers.HandleDeleteProduct)

	front.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
