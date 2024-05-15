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

const (
	defaultHost     = "database_postgres"
	defaultPort     = "5432"
	defaultUser     = "postgres"
	defaultPassword = "mindmentor"
	defaultDBName   = "mindmentor"
)

func corsMiddleware(next http.Handler) http.Handle {
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

	// Инициализация репозитория пользователей
	userRepo := repositories.NewUserRepository(db)

	// Инициализация сервиса авторизации
	authService := services.NewAuthService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)

	// Настройка HTTP обработчиков
	http.Handle("/register", corsMiddleware(http.HandlerFunc(authHandler.RegisterUserHandler)))
	http.Handle("/login", corsMiddleware(http.HandlerFunc(authHandler.AuthenticateUserHandler)))

	// Запуск сервера
	fmt.Println("Authentication service is running on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", defaultHost, defaultPort, defaultUser, defaultPassword, defaultDBName)
	db, err := sql.Open("postgres", connStr)

	return db, err
}
