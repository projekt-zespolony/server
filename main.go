package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/sensors", handleGetSensors)
	e.POST("/sensors", handlePostSensors, middleware.KeyAuth(handleAuth))
	log.Fatal(e.Start(":8080"))
}
