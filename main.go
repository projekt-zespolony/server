package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)


func handleTest (w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "response")
}


func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", handleTest).Methods(http.MethodGet)
    server := http.Server{
        Addr:              "0.0.0.0:8080",
        Handler:           router,
        TLSConfig:         nil,
    }
    err := server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}
