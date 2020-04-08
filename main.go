package main

import (
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

var (
	serverPort = os.Getenv("SERVER_PORT")
	accessToken = os.Getenv("SERVER_ACCESS_TOKEN")
	certsCacheDir = os.Getenv("SERVER_CERTS_CACHE_DIR")

	currentSensors = &Sensors{
		Timestamp: time.Now().Unix(),
	}
)

func main() {
	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache(certsCacheDir)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/sensors", handleGetSensors)
	e.POST("/sensors", handlePostSensors, middleware.KeyAuth(handleAuth))
	e.POST("/firebase", handleFirebase, middleware.KeyAuth(handleAuth))
	log.Fatal(e.StartAutoTLS(serverPort))
}
