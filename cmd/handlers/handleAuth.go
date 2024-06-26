package handlers

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo/v4"
)

const SECRET = "super_secret_key"

type JwtClaims struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	credentials := models.Credentials{}
	c.Bind(&credentials)

	/*if credentials.Username != "walter" && credentials.Password != "secret" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Wrong credentials", "message": "Access denied"})
	}*/

	user, isCorrect := repositories.AuthenticateUser(credentials.Username, credentials.Password)
	if !isCorrect {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Wrong credentials", "message": "Access denied"})
	}

	claims := &JwtClaims{
		Name:  user.Name,
		UUID:  "444-555-dscedf-kaqwsa",
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	fmt.Println("Claims", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"token": t})
}
