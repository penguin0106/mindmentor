package main

import (
	"database/sql"
	"log"
	"mindmentor/services/emotions_service/handlers"
	"mindmentor/services/emotions_service/repositories"
	"mindmentor/services/emotions_service/services"
	"net/http"
)

func main() {
	// Подключение к базе данных
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Инициализация репозитория
	emotionRepo := repositories.NewEmotionRepository(db)

	emoService := services.NewEmotionService(emotionRepo)

	// Инициализация обработчика
	emotionHandler := handlers.NewEmotionHandler(emoService)

	// Регистрация HTTP обработчиков
	http.HandleFunc("/emotions", emotionHandler.CreateEmotionHandler)
	http.HandleFunc("/emotions/update", emotionHandler.UpdateEmotionHandler)
	http.HandleFunc("/emotions/delete", emotionHandler.DeleteEmotionHandler)
	http.HandleFunc("/emotions/user", emotionHandler.GetEmotionsByUserHandler)

	// Запуск сервера
	http.ListenAndServe(":8082", nil)
}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@localhost:5432/mindmentor?sslmode=disable")
	return db, err
}
