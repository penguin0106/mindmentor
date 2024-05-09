package main

import (
	"database/sql"
	"emotions_service/handlers"
	"emotions_service/repositories"
	"emotions_service/services"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			http.Error(w, "", http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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

	http.Handle("/emotions", corsMiddleware(http.HandlerFunc(emotionHandler.CreateEmotionHandler)))
	http.Handle("/emotions/update", corsMiddleware(http.HandlerFunc(emotionHandler.UpdateEmotionHandler)))
	http.Handle("/emotions/delete", corsMiddleware(http.HandlerFunc(emotionHandler.DeleteEmotionHandler)))
	http.Handle("/emotions/user", corsMiddleware(http.HandlerFunc(emotionHandler.GetEmotionsByUserHandler)))

	// Запуск сервера
	http.ListenAndServe(":8082", nil)
}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@localhost:5432/mindmentor?sslmode=disable")
	return db, err
}
