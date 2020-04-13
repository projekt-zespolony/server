package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/projekt-zespolony/server/pkg/database"
	"github.com/projekt-zespolony/server/pkg/firebase"
	"github.com/projekt-zespolony/server/pkg/types"
	"golang.org/x/crypto/acme/autocert"
)

type Router struct {
	fb      *firebase.Firebase
	db      *database.Database
	options *Options
}

type Options struct {
	Version       string
	Commit        string
	AccessToken   string
	ServerPort    string
	CertsCacheDir string
}

func Run(routerOptions *Options, dbOptions *database.Options, firebaseOptions *firebase.Options) error {
	fb, err := firebase.New(firebaseOptions)
	if err != nil {
		return err
	}

	db, err := database.New(dbOptions)
	if err != nil {
		return err
	}

	router := &Router{
		fb:      fb,
		db:      db,
		options: routerOptions,
	}

	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache(routerOptions.CertsCacheDir)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", router.handleGetStatus)
	e.GET("/sensors", router.handleGetSensors)
	e.POST("/sensors", router.handlePostSensors, middleware.KeyAuth(router.handleAuth))
	e.POST("/firebase", router.handlePostFirebase, middleware.KeyAuth(router.handleAuth))

	if routerOptions.ServerPort == "443" {
		return e.StartAutoTLS(":" + routerOptions.ServerPort)
	} else {
		return e.Start(":" + routerOptions.ServerPort)
	}
}

func (router *Router) handleAuth(key string, c echo.Context) (bool, error) {
	if key == router.options.AccessToken {
		return true, nil
	}
	return false, nil
}

func (router *Router) handleGetStatus(c echo.Context) error {
	status := &types.Status{
		Version: router.options.Version,
		Commit:  router.options.Commit,
	}

	return c.JSON(http.StatusOK, status)
}

func (router *Router) handleGetSensors(c echo.Context) error {
	sensors, err := router.db.Latest()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, sensors)
}

func (router *Router) handlePostSensors(c echo.Context) error {
	sensors := &types.Sensors{
		Timestamp: time.Now().UTC().Unix(),
	}

	err := c.Bind(sensors)
	if err != nil {
		return err
	}

	err = router.db.Create(sensors)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, sensors)
}

func (router *Router) handlePostFirebase(c echo.Context) error {
	notification := &types.Notification{
		Topic: "default",
	}

	err := c.Bind(notification)
	if err != nil {
		return err
	}

	err = router.fb.Notify(notification)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, notification)
}
