package main

import (
	"auth_service/handlers"
	"auth_service/repositories"
	"auth_service/services"
	"database/sql"
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

	// Инициализация репозитория пользователей
	userRepo := repositories.NewUserRepository(db)

	// Инициализация сервиса авторизации
	authService := services.NewAuthService(userRepo)

	// Настройка HTTP обработчиков
	http.HandleFunc("/register", handlers.RegisterHandler(authService))
	http.HandleFunc("/login", handlers.LoginHandler(authService))

	// Запуск сервера
	fmt.Println("Authentication service is running on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@localhost:5432/mindmentor?sslmode=disable")
	return db, err
}
