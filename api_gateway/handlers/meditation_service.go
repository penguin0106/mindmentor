package handlers

import "net/http"

const meditationServiceURL = "http://meditation_service:8083"

// MedCoursesAllHandler КУРСЫ
func MedCoursesAllHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/courses/all", r.Method, r.Body)
}

func MedCoursesSearchHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/course/search", r.Method, r.Body)
}

func MedCoursesAddHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the meditation service
	proxyRequest(w, meditationServiceURL+"/courses/add", r.Method, r.Body)
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
