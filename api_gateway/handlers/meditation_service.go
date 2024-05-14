package handlers

import "net/http"

const meditationServiceURL = "http://meditation_service:8083"

// MedVideoAllHandler КУРСЫ
func MedVideoAllHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/video/all", r.Method, r.Body)
}

func MedVideoSearchHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/video/search", r.Method, r.Body)
}

func MedVideoAddHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/video/add", r.Method, r.Body)
}

func MedVideoDeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/video/delete", r.Method, r.Body)
}

// MedMusicAllHandler МУЗЫКА
func MedMusicAllHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/music/all", r.Method, r.Body)
}

func MedMusicAddHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/music/add", r.Method, r.Body)
}

// MedRatingsGetHandler Рейтинг
func MedRatingsGetHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/ratings/get", r.Method, r.Body)
}

func MedRatingsAddHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/ratings/add", r.Method, r.Body)
}

// MedCommentsGetHandler Comments
func MedCommentsGetHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/comments/get", r.Method, r.Body)
}

func MedCommentsAddHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/comments/add", r.Method, r.Body)
}

// MedFavouritesAddHandler Favourites
func MedFavouritesAddHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/favorites/add", r.Method, r.Body)
}

func MedFavouritesRemoveHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/favorites/remove", r.Method, r.Body)
}
