package main

import (
	"log"
	"os"

	"github.com/projekt-zespolony/server/pkg/database"
	"github.com/projekt-zespolony/server/pkg/server"
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

	serverOptions := &server.Options{
		Version:       "",
		Commit:        "",
		AccessToken:   os.Getenv("SERVER_ACCESS_TOKEN"),
		ServerPort:    os.Getenv("SERVER_PORT"),
		CertsCacheDir: os.Getenv("SERVER_CERTS_CACHE_DIR"),
	}

	log.Fatal(server.Run(serverOptions, dbOptions))
}
