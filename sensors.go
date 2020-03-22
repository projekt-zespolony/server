package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Sensors struct {
	Timestamp   int64   `json:"timestamp"`
	Temperature float32 `json:"temperature"`
	Humidity    int32   `json:"humidity"`
}

var currentSensors = &Sensors{
	Timestamp:   time.Now().Unix(),
	Temperature: 0,
	Humidity:    0,
}

func handleGetSensors(writer http.ResponseWriter, request *http.Request) {
	body, err := json.MarshalIndent(currentSensors, "", "  ")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(body)
}

func handlePostSensors(writer http.ResponseWriter, request *http.Request) {
	if !isAuthorized(request) {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, currentSensors)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err = json.MarshalIndent(currentSensors, "", "  ")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(body)
}
