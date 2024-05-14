package handlers

import (
	"net/http"
)

const (
	authServiceURL = "http://auth_service:8081"
)

func AuthRegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the auth service
	proxyRequest(w, authServiceURL+"/register", r.Method, r.Body)
}

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the auth service
	proxyRequest(w, authServiceURL+"/login", r.Method, r.Body)
}
