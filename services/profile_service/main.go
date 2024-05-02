package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"mindmentor/services/profile_service/handlers"
	"mindmentor/services/profile_service/repositories"
	"mindmentor/services/profile_service/services"
	"net/http"
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
	http.HandleFunc("/user", userHandler.GetUserHandler)
	http.HandleFunc("/user/update", userHandler.UpdateUserHandler)
	http.HandleFunc("/favorites/course-get", userHandler.GetFavoriteCourseHandler)
	http.HandleFunc("/favorites/training-get", userHandler.GetFavoriteTrainingHandler)

	// Запуск сервера
	log.Println("Server started on port 8086")
	log.Fatal(http.ListenAndServe(":8086", nil))
}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@localhost:5432/mindmentor?sslmode=disable")
	return db, err
}
