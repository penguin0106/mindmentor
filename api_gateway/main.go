package main

import (
	"fmt"
	"io"
	"log"
	"mindmentor/api_gateway/middleware"
	"net/http"
)

const (
	databaseServiceURL   = "http://localhost:5432"
	authServiceURL       = "http://localhost:8081"
	emotionsServiceURL   = "http://localhost:8082"
	meditationServiceURL = "http://localhost:8083"
	socialServiceURL     = "http://localhost:8084"
	trainingsServiceURL  = "http://localhost:8085"
	profileServiceURL    = "http://localhost:8086"
)

func main() {
	// Set up middleware
	authMiddleware := middleware.AuthMiddleware
	loggingMiddleware := middleware.LoggingMiddleware

	// Apply middleware to handlers
	http.HandleFunc("/auth", loggingMiddleware(authMiddleware(middleware.WrapHandlerFunc(authHandler))))
	http.HandleFunc("/database", loggingMiddleware(authMiddleware(middleware.WrapHandlerFunc(databaseHandler))))
	http.HandleFunc("/emotions", loggingMiddleware(authMiddleware(middleware.WrapHandlerFunc(emotionsHandler))))
	http.HandleFunc("/meditation", loggingMiddleware(authMiddleware(middleware.WrapHandlerFunc(meditationHandler))))
	http.HandleFunc("/profile", loggingMiddleware(authMiddleware(middleware.WrapHandlerFunc(profileHandler))))
	http.HandleFunc("/social", loggingMiddleware(authMiddleware(middleware.WrapHandlerFunc(socialHandler))))
	http.HandleFunc("/trainings", loggingMiddleware(authMiddleware(middleware.WrapHandlerFunc(trainingsHandler))))

	fmt.Println("API Gateway is running on port 8090...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the auth service
	proxyRequest(w, authServiceURL+r.URL.Path)
}

func databaseHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the database service
	proxyRequest(w, databaseServiceURL+r.URL.Path)
}

func emotionsHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the emotions service
	proxyRequest(w, emotionsServiceURL+r.URL.Path)
}

func meditationHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+r.URL.Path)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+r.URL.Path)
}

func socialHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the social service
	proxyRequest(w, socialServiceURL+r.URL.Path)
}

func trainingsHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the trainings service
	proxyRequest(w, trainingsServiceURL+r.URL.Path)
}

func proxyRequest(w http.ResponseWriter, url string) {
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to proxy request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
