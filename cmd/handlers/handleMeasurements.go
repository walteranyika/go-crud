package handlers

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateMeasurement(c echo.Context) error{
    measurement := models.Measurement{}
	c.Bind(&measurement)
	newMeasurement , err := repositories.CreateMeasurement(measurement)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, newMeasurement)
}

func UpdateMeasurement(c echo.Context) error{
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err!=nil {
		return c.JSON(http.StatusBadRequest, err.Error(),)
	}

	measurement := models.Measurement{}
	c.Bind(&measurement)

	updatedMeasurement, err := repositories.UpdateMeasurement(measurement, intId)
	if err!=nil {
		return c.JSON(http.StatusBadRequest, err.Error(),)
	}
	return c.JSON(http.StatusAccepted, updatedMeasurement)
}