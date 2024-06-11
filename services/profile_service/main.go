package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"profile_service/handlers"
	"profile_service/repositories"
	"profile_service/services"
)

const (
	defaultHost     = "database_postgres"
	defaultPort     = "5432"
	defaultUser     = "postgres"
	defaultPassword = "mindmentor"
	defaultDBName   = "mindmentor"
)

func main() {
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Инициализация репозиториев
	userRepository := repositories.NewUserRepository(db)

	// Инициализация сервисов
	userService := services.NewUserService(userRepository)

	// Инициализация обработчиков запросов
	userHandler := handlers.NewUserHandler(userService)

	// Настройка маршрутов
	http.Handle("/user/get", corsMiddleware(http.HandlerFunc(userHandler.GetUserHandler)))
	http.Handle("/user/change_username", corsMiddleware(http.HandlerFunc(userHandler.EditProfileUsernameHandler)))
	http.Handle("/user/change_email", corsMiddleware(http.HandlerFunc(userHandler.EditProfileEmailHandler)))
	http.Handle("/user/change_password", corsMiddleware(http.HandlerFunc(userHandler.EditProfilePasswordHandler)))
	http.Handle("/favorites/video", corsMiddleware(http.HandlerFunc(userHandler.GetFavoriteVideosHandler)))
	http.Handle("/favorites/training", corsMiddleware(http.HandlerFunc(userHandler.GetFavoriteTrainingsHandler)))

	// Запуск сервера
	log.Println("Server started on port 8086")
	log.Fatal(http.ListenAndServe(":8086", nil))
}
