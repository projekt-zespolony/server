package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sensors", handleGetSensors).Methods(http.MethodGet)
	router.Handle("/sensors", NeedsAuth(handlePostSensors)).Methods(http.MethodPost)
	server := http.Server{
		Addr:      "0.0.0.0:8080",
		Handler:   router,
		TLSConfig: nil,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
