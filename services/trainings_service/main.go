package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"trainings_service/handlers"
	"trainings_service/repositories"
	"trainings_service/services"
)

const (
	defaultHost     = "database_postgres"
	defaultPort     = "5432"
	defaultUser     = "postgres"
	defaultPassword = "mindmentor"
	defaultDBName   = "mindmentor"
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
	http.Handle("/trainings/get", corsMiddleware(http.HandlerFunc(trainingHandler.GetAllBooksHandler)))
	http.Handle("/trainings/search", corsMiddleware(http.HandlerFunc(trainingHandler.GetBookByNameHandler)))

	// Register favorite handler functions with CORS middleware
	http.Handle("/favorites/add", corsMiddleware(http.HandlerFunc(favoriteHandler.AddToFavoritesHandler)))
	http.Handle("/favorites/remove", corsMiddleware(http.HandlerFunc(favoriteHandler.RemoveFromFavoritesHandler)))

	// Register comment handler functions with CORS middleware
	http.Handle("/comments/add", corsMiddleware(http.HandlerFunc(commentHandler.AddBookCommentHandler)))
	http.Handle("/comments/get", corsMiddleware(http.HandlerFunc(commentHandler.GetBookCommentsHandler)))

	http.Handle("/rating/add", corsMiddleware(http.HandlerFunc(commentHandler.AddBookRatingHandler)))
	http.Handle("/rating/get", corsMiddleware(http.HandlerFunc(commentHandler.GetBookRatingHandler)))

	// Запуск сервера
	http.ListenAndServe(":8085", nil)
}
