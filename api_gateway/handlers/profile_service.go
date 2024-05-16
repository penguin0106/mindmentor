package handlers

import "net/http"

const profileServiceURL = "http://profile_service:8086"

func ProfileUserGetHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+"/user/get", r.Method, r.Body)
}

func ProfileEditProfileUsernameHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+"/user/change_username", r.Method, r.Body)
}

func ProfileEditProfileEmailHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+"/user/change_email", r.Method, r.Body)
}

func ProfileEditProfilePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+"/user/change_password", r.Method, r.Body)
}

func ProfileFavouritesCoursesHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+"/favorites/video", r.Method, r.Body)
}

func ProfileFavouritesTrainingsHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the profile service
	proxyRequest(w, profileServiceURL+"/favorites/training", r.Method, r.Body)
}
