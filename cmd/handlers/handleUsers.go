package handlers

import (
	"errors"
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

// https://www.kelche.co/blog/go/golang-echo-tutorial/
func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	errMap, er := validateUser(&user)
	if er != nil {
		return c.JSON(http.StatusBadRequest, errMap)
	}

	newUser, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := models.User{}
	c.Bind(&user)
	updatedUser, err := repositories.UpdateUser(user, idInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, updatedUser)
}

func GetAllUsers(c echo.Context) error {
	users, err := repositories.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := repositories.GetUser(idInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func validateUser(user *models.User) (map[string]interface{}, error) {
	rules := govalidator.MapData{
		"name":     []string{"required", "between:5,25"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"password": []string{"required", "min:6", "max:30"},
	}

	options := govalidator.Options{
		Data:  user,
		Rules: rules,
	}
	validator := govalidator.New(options)
	e := validator.ValidateStruct()
	if e != nil {
	 erroMap := map[string]interface{}{"validationError": e}
	 return erroMap, errors.New("invalid")
	}
    return map[string]interface{}{"validationError": nil}, nil
}
