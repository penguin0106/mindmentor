package main

import (
	"database/sql"
	"log"
	"mindmentor/services/trainings_service/handlers"
	"mindmentor/services/trainings_service/repositories"
	"mindmentor/services/trainings_service/services"
	"net/http"
)

func main() {
	// Подключение к базе данных
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Инициализация репозитория
	trainingRepo := repositories.NewTrainingRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	favouriteRepo := repositories.NewFavoriteRepository(db)

	trainingServ := services.NewTrainingService(trainingRepo)
	commentServ := services.NewCommentService(commentRepo)
	favoriteServ := services.NewFavoriteService(favouriteRepo)

	// Инициализация обработчика избранных тренировок
	trainingHandler := handlers.NewTrainingHandler(trainingServ)
	commentHandler := handlers.NewCommentHandler(commentServ)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteServ)

	// Регистрация HTTP обработчиков
	http.HandleFunc("/trainings", trainingHandler.GetAllTrainingsHandler)
	http.HandleFunc("/trainings/search", trainingHandler.GetTrainingByNameHandler)
	http.HandleFunc("/favorites/add", favoriteHandler.AddToFavoritesHandler)
	http.HandleFunc("/favorites/remove", favoriteHandler.RemoveFromFavoritesHandler)
	http.HandleFunc("/comments", commentHandler.AddCommentHandler)
	http.HandleFunc("/comments/trainings", commentHandler.GetCommentsByTrainingIDHandler)
	http.HandleFunc("/rating/add", commentHandler.AddRatingHandler)
	http.HandleFunc("/rating/get", commentHandler.GetRatingHandler)

	// Запуск сервера
	http.ListenAndServe(":8085", nil)
}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@localhost:5432/mindmentor?sslmode=disable")
	return db, err
}
