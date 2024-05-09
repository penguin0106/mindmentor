package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"profile_service/handlers"
	"profile_service/repositories"
	"profile_service/services"
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
	http.Handle("/user", corsMiddleware(http.HandlerFunc(userHandler.GetUserHandler)))
	http.Handle("/user/update", corsMiddleware(http.HandlerFunc(userHandler.UpdateUserHandler)))
	http.Handle("/favorites/course-get", corsMiddleware(http.HandlerFunc(userHandler.GetFavoriteCourseHandler)))
	http.Handle("/favorites/training-get", corsMiddleware(http.HandlerFunc(userHandler.GetFavoriteTrainingHandler)))

	// Запуск сервера
	log.Println("Server started on port 8086")
	log.Fatal(http.ListenAndServe(":8086", nil))
}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@localhost:5432/mindmentor?sslmode=disable")
	return db, err
}
