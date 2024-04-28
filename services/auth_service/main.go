package main

import (
	"database/sql"
	"fmt"
	"log"
	"mindmentor/services/auth_service/handlers"
	"mindmentor/services/auth_service/repositories"
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
	db, err := sql.Open("postgres", "postgres://tdteam:tdteam@localhost:5432/tdteam?sslmode=disable")
	return db, err
}
