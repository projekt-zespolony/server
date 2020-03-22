package main

import (
	"net/http"
)

var token = "@@TOKEN@@"

func isAuthorized(request *http.Request) bool {
	if request.Header.Get("token") != token {
		return false
	}

	return true
}
