package main

import (
	"log"
	"os"

	"github.com/projekt-zespolony/server/pkg/database"
	"github.com/projekt-zespolony/server/pkg/router"
)

var (
	version string
	commit  string
)

func main() {
	dbOptions := &database.Options{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
		Addr: os.Getenv("DB_ADDR"),
	}

	routerOptions := &router.Options{
		Version:       version,
		Commit:        commit,
		AccessToken:   os.Getenv("SERVER_ACCESS_TOKEN"),
		ServerPort:    os.Getenv("SERVER_PORT"),
		CertsCacheDir: os.Getenv("SERVER_CERTS_CACHE_DIR"),
	}

	log.Fatal(router.Run(routerOptions, dbOptions))
}
