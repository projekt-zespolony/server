package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Sensors struct {
	Timestamp   int64   `json:"timestamp"`
	Temperature float32 `json:"temperature"`
    Pressure    float32 `json:"pressure"`
	Humidity    float32 `json:"humidity"`
    Gas         float32 `json:"gas"`
}

var currentSensors = &Sensors{
	Timestamp:   time.Now().Unix(),
}

func handleGetSensors(c echo.Context) error {
	return c.JSON(http.StatusOK, currentSensors)
}

func handlePostSensors(c echo.Context) error {
	err := c.Bind(currentSensors)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, currentSensors)
}
