package main

import (
	"database/sql"
	"emotions_service/handlers"
	"emotions_service/repositories"
	"emotions_service/services"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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

	http.Handle("/create", corsMiddleware(http.HandlerFunc(emotionHandler.CreateEmotionHandler)))
	http.Handle("/update", corsMiddleware(http.HandlerFunc(emotionHandler.UpdateEmotionHandler)))
	http.Handle("/delete", corsMiddleware(http.HandlerFunc(emotionHandler.DeleteEmotionHandler)))
	http.Handle("/user", corsMiddleware(http.HandlerFunc(emotionHandler.GetEmotionsByUserHandler)))

	// Запуск сервера
	http.ListenAndServe(":8082", nil)
}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", defaultHost, defaultPort, defaultUser, defaultPassword, defaultDBName)
	db, err := sql.Open("postgres", connStr)

	return db, err
}
