package main

import (
	"fitness-api/cmd/handlers"
	"fitness-api/cmd/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(handlers.LogRequest)
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: [] string{ echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
    e.GET("/", handlers.Home)

	storage.InitDB()

	e.POST("/users", handlers.CreateUser)
	e.POST("/measurements", handlers.CreateMeasurement)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.PUT("/measurements/:id", handlers.UpdateMeasurement)
	e.Logger.Fatal(e.Start(":3000"))
}