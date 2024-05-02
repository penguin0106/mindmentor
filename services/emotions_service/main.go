package main

import (
	"mindmentor/services/emotions_service/handlers"
	"mindmentor/services/emotions_service/repositories"
	"net/http"
)

func main() {
	// Инициализация репозитория
	emotionRepo := &repositories.EmotionRepository{}

	// Инициализация обработчика
	emotionHandler := &handlers.EmotionHandler{Repository: emotionRepo}

	// Регистрация HTTP обработчиков
	http.HandleFunc("/emotions", emotionHandler.CreateEmotionHandler)
	http.HandleFunc("/emotions/update", emotionHandler.UpdateEmotionHandler)
	http.HandleFunc("/emotions/delete", emotionHandler.DeleteEmotionHandler)
	http.HandleFunc("/emotions/user", emotionHandler.GetEmotionsByUserHandler)

	// Запуск сервера
	http.ListenAndServe(":8082", nil)
}
