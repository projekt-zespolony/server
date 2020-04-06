package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	token = "@@TOKEN@@"

	currentSensors = &Sensors{
		Timestamp: time.Now().Unix(),
	}
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/sensors", handleGetSensors)
	e.POST("/sensors", handlePostSensors, middleware.KeyAuth(handleAuth))
	e.POST("/firebase", handleFirebase, middleware.KeyAuth(handleAuth))
	log.Fatal(e.Start(":8080"))
}
