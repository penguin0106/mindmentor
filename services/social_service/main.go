package main

import (
	"mindmentor/services/social_service/handlers"
	"mindmentor/services/social_service/repositories"
	"net/http"
)

func main() {
	// Инициализация репозиториев
	discussionRepo := &repositories.DiscussionRepository{}

	// Инициализация обработчиков
	discussionHandler := &handlers.DiscussionHandler{Repo: discussionRepo}

	// Регистрация HTTP обработчиков
	http.HandleFunc("/discussions/create", discussionHandler.CreateDiscussionHandler)
	http.HandleFunc("/discussions/find", discussionHandler.FindDiscussionHandler)
	http.HandleFunc("/discussions/join", discussionHandler.JoinDiscussionHandler)
	http.HandleFunc("/discussions/leave", discussionHandler.LeaveDiscussionHandler)

	// Запуск сервера
	http.ListenAndServe(":8080", nil)
}
