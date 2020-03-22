package main

import (
	"net/http"
)

var token = "@@TOKEN@@"

type Auth struct {
	http.Handler
	handlerFunc http.HandlerFunc
}

func (auth *Auth) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("token") != token {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	auth.handlerFunc(writer, request)
}

func NeedsAuth(handlerFunc http.HandlerFunc) http.Handler {
	return &Auth{
		handlerFunc: handlerFunc,
	}
}
