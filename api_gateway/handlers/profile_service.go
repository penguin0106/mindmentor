package handlers

import "net/http"

const profileServiceURL = "http://profile_service:8086"

func ProfileUserGetHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+"/user/get", r.Method, r.Body)
}

func ProfileUserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+"/user/update", r.Method, r.Body)
}

func ProfileFavouritesCoursesHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+"/favorites/course", r.Method, r.Body)
}

func ProfileFavouritesTrainingsHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+"/favorites/training", r.Method, r.Body)
}
