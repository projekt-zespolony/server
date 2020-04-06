package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleAuth(key string, c echo.Context) (bool, error) {
	if key == token {
		return true, nil
	}
	return false, nil
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

func handleFirebase(c echo.Context) error {
	notification := &Notification{}

	err := c.Bind(notification)
	if err != nil {
		return err
	}

	err = firebaseNotify(notification)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, notification)
}
