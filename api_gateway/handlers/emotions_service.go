package handlers

import "net/http"

const emotionsServiceURL = "http://emotions_service:8082"

func EmotionsCreateHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the emotions service
	proxyRequest(w, emotionsServiceURL+"/create", r.Method, r.Body)
}

func EmotionsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the emotions service
	proxyRequest(w, emotionsServiceURL+"/update", r.Method, r.Body)
}

func EmotionsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the emotions service
	proxyRequest(w, emotionsServiceURL+"/delete", r.Method, r.Body)
}

func EmotionsUserHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the emotions service
	proxyRequest(w, emotionsServiceURL+"/user", r.Method, r.Body)
}
