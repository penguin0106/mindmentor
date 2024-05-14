package handlers

import "net/http"

const socialServiceURL = "http://social_service:8084"

// SocialDiscussionAddHandler Discussions
func SocialDiscussionAddHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the social service
	proxyRequest(w, socialServiceURL+"/discussions/add", r.Method, r.Body)
}

func SocialDiscussionFindHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the social service
	proxyRequest(w, socialServiceURL+"/discussions/find", r.Method, r.Body)
}

func SocialDiscussionJoinHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the social service
	proxyRequest(w, socialServiceURL+"/discussions/join", r.Method, r.Body)
}

func SocialDiscussionLeaveHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the social service
	proxyRequest(w, socialServiceURL+"/discussions/leave", r.Method, r.Body)
}

// SocialMessageSendHandler Messages
func SocialMessageSendHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the social service
	proxyRequest(w, socialServiceURL+"/messages/send", r.Method, r.Body)
}

func SocialMessageEditHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the social service
	proxyRequest(w, socialServiceURL+"/messages/edit", r.Method, r.Body)
}

func SocialMessageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Proxy the request to the social service
	proxyRequest(w, socialServiceURL+"/messages/delete", r.Method, r.Body)
}
