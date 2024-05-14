package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"social_service/handlers"
	"social_service/repositories"
	"social_service/services"
)

const (
	defaultHost     = "database_postgres"
	defaultPort     = "5432"
	defaultUser     = "postgres"
	defaultPassword = "mindmentor"
	defaultDBName   = "mindmentor"
)

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", defaultHost, defaultPort, defaultUser, defaultPassword, defaultDBName)
	db, err := sql.Open("postgres", connStr)

	return db, err
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, UPDATE, DELETE, PUT, OPTIONS")
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
	// Инициализация репозиториев
	discussionRepo := repositories.NewDiscussionRepository(db)
	messageRepo := repositories.NewMessageRepository(db)

	discService := services.NewDiscussionService(discussionRepo)
	messageService := services.NewMessageService(messageRepo)

	// Инициализация обработчиков
	discussionHandler := handlers.NewDiscussionHandler(discService)
	messageHandler := handlers.NewMessageHandler(messageService)

	// Регистрация HTTP обработчиков

	// Register discussion handler functions with CORS middleware
	http.Handle("/discussions/add", corsMiddleware(http.HandlerFunc(discussionHandler.AddDiscussionHandler)))
	http.Handle("/discussions/find", corsMiddleware(http.HandlerFunc(discussionHandler.FindDiscussionHandler)))
	http.Handle("/discussions/join", corsMiddleware(http.HandlerFunc(discussionHandler.JoinDiscussionHandler)))
	http.Handle("/discussions/leave", corsMiddleware(http.HandlerFunc(discussionHandler.LeaveDiscussionHandler)))

	// Register message handler functions with CORS middleware
	http.Handle("/messages/send", corsMiddleware(http.HandlerFunc(messageHandler.SendMessageHandler)))
	http.Handle("/messages/edit", corsMiddleware(http.HandlerFunc(messageHandler.EditMessageHandler)))
	http.Handle("/messages/delete", corsMiddleware(http.HandlerFunc(messageHandler.DeleteMessageHandler)))

	// Запуск сервера
	http.ListenAndServe(":8084", nil)
}
