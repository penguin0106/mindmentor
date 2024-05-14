package handlers

import "net/http"

const trainingsServiceURL = "http://trainings_service:8085"

func TrainingsGetHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the trainings service
	proxyRequest(w, trainingsServiceURL+"/trainings/get", r.Method, r.Body)
}

func TrainingsSearchHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the trainings service
	proxyRequest(w, trainingsServiceURL+"/trainings/search", r.Method, r.Body)
}

func TrainingsFavouritesAddHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the trainings service
	proxyRequest(w, trainingsServiceURL+"/favorites/add", r.Method, r.Body)
}

func TrainingsFavouritesRemoveHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the trainings service
	proxyRequest(w, trainingsServiceURL+"/favorites/remove", r.Method, r.Body)
}

func TrainingsCommentsAddHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the trainings service
	proxyRequest(w, trainingsServiceURL+"/comments/add", r.Method, r.Body)
}

func TrainingsCommentsGetHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the trainings service
	proxyRequest(w, trainingsServiceURL+"/comments/get", r.Method, r.Body)
}

func TrainingsRatingAddHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the trainings service
	proxyRequest(w, trainingsServiceURL+"/rating/add", r.Method, r.Body)
}

func TrainingsRatingGetHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the trainings service
	proxyRequest(w, trainingsServiceURL+"/rating/get", r.Method, r.Body)
}
