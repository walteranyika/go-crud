package main

import (
	"fitness-api/cmd/handlers"
	"fitness-api/cmd/storage"
	"net/http"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(handlers.LogRequest)
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", handlers.Home)

	storage.InitDB()

	e.POST("/login", handlers.Login)

	securedGroup := e.Group("/api")
	config := middleware.JWTConfig{
		Claims:     &handlers.JwtClaims{},
		SigningKey: []byte(handlers.SECRET),
	}
	securedGroup.Use(middleware.JWTWithConfig(config))
	securedGroup.GET("/test/user", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*handlers.JwtClaims)
		name := claims.Name
		return c.JSON(http.StatusOK, map[string]string{"name": name})
	})

	securedGroup.GET("/users", handlers.GetAllUsers)
	securedGroup.GET("/users/:id", handlers.GetUser)
	securedGroup.POST("/users", handlers.CreateUser)
	securedGroup.POST("/measurements", handlers.CreateMeasurement)
	securedGroup.PUT("/users/:id", handlers.UpdateUser)
	securedGroup.PUT("/measurements/:id", handlers.UpdateMeasurement)

	e.Logger.Fatal(e.Start(":3000"))
}
