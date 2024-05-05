package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"social_service/handlers"
	"social_service/repositories"
	"social_service/services"
)

func main() {
	// Подключение к базе данных
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Инициализация репозиториев
	discussionRepo := repositories.NewDiscussionRepository(db)

	discService := services.NewDiscussionService(discussionRepo)

	// Инициализация обработчиков
	discussionHandler := handlers.NewDiscussionHandler(discService)

	// Регистрация HTTP обработчиков
	http.HandleFunc("/discussions/create", discussionHandler.CreateDiscussionHandler)
	http.HandleFunc("/discussions/find", discussionHandler.FindDiscussionHandler)
	http.HandleFunc("/discussions/join", discussionHandler.JoinDiscussionHandler)
	http.HandleFunc("/discussions/leave", discussionHandler.LeaveDiscussionHandler)
	http.HandleFunc("/messages/update", discussionHandler.UpdateMessageHandler)
	http.HandleFunc("/messages/delete", discussionHandler.DeleteMessageHandler)

	// Запуск сервера
	http.ListenAndServe(":8084", nil)
}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@localhost:5432/mindmentor?sslmode=disable")
	return db, err
}
