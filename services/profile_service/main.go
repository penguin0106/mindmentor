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
