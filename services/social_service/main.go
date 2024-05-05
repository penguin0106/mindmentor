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
	messageRepo := repositories.NewMessageRepository(db)

	discService := services.NewDiscussionService(discussionRepo)
	messageService := services.NewMessageService(messageRepo)

	// Инициализация обработчиков
	discussionHandler := handlers.NewDiscussionHandler(discService)
	messageHandler := handlers.NewMessageHandler(messageService)

	// Регистрация HTTP обработчиков
	http.HandleFunc("/discussions/add", discussionHandler.AddDiscussionHandler)
	http.HandleFunc("/discussions/find", discussionHandler.FindDiscussionHandler)
	http.HandleFunc("/discussions/join", discussionHandler.JoinDiscussionHandler)
	http.HandleFunc("/discussions/leave", discussionHandler.LeaveDiscussionHandler)
	http.HandleFunc("messages/send", messageHandler.SendMessageHandler)
	http.HandleFunc("/messages/edit", messageHandler.EditMessageHandler)
	http.HandleFunc("/messages/delete", messageHandler.DeleteMessageHandler)

	// Запуск сервера
	http.ListenAndServe(":8084", nil)
}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@localhost:5432/mindmentor?sslmode=disable")
	return db, err
}
