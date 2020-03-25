package main

import (
	"github.com/labstack/echo/v4"
)

var token = "TOKEN"

func handleAuth(key string, c echo.Context) (bool, error) {
	if key == token {
		return true, nil
	}
	return false, nil
}
