package main

import (
	"fitness-api/cmd/handlers"
	"fitness-api/cmd/storage"
	"net/http"

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


	e.POST("/login", handlers.Login)

	protected := e.Group("/api")
	config := middleware.JWTConfig{
		Claims: &handlers.JwtClaims{},
		SigningKey: []byte(handlers.SECRET),
	}
    protected.Use(middleware.JWTWithConfig(config))
	protected.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"name":"Walter"})
	})

	protected.POST("/users", handlers.CreateUser)
	protected.POST("/measurements", handlers.CreateMeasurement)
	protected.PUT("/users/:id", handlers.UpdateUser)
	protected.PUT("/measurements/:id", handlers.UpdateMeasurement)
	
	e.Logger.Fatal(e.Start(":3000"))
}