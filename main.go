package main

import (
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	token = os.Getenv("TOKEN")

	currentSensors = &Sensors{
		Timestamp: time.Now().Unix(),
	}
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/sensors", handleGetSensors)
	e.POST("/sensors", handlePostSensors, middleware.KeyAuth(handleAuth))
	e.POST("/firebase", handleFirebase, middleware.KeyAuth(handleAuth))
	log.Fatal(e.Start(":8080"))
}
