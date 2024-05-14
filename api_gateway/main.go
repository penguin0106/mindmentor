package main

import (
	"api_gateway/middleware"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
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
	http.HandleFunc("/auth", authHandler)
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
	proxyRequest(w, authServiceURL+r.URL.Path, r.Method, r.Body)
}

func emotionsHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the emotions service
	proxyRequest(w, emotionsServiceURL+r.URL.Path, r.Method, r.Body)
}

func meditationHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+r.URL.Path, r.Method, r.Body)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+r.URL.Path, r.Method, r.Body)
}

func socialHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the social service
	proxyRequest(w, socialServiceURL+r.URL.Path, r.Method, r.Body)
}

func trainingsHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the trainings service
	proxyRequest(w, trainingsServiceURL+r.URL.Path, r.Method, r.Body)
}

func proxyRequest(w http.ResponseWriter, url string, method string, body io.Reader) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to proxy request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Копирование заголовков ответа
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
