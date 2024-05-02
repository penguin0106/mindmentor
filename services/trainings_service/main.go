package main

import (
	"mindmentor/services/trainings_service/handlers"
	"mindmentor/services/trainings_service/repositories"
	"net/http"
)

func main() {
	// Инициализация репозитория
	trainingRepo := &repositories.TrainingRepository{}
	commentRepo := &repositories.CommentRepository{}
	favouriteRepo := &repositories.FavoriteRepository{}

	// Инициализация обработчика избранных тренировок
	trainingHandler := &handlers.TrainingHandler{Repository: trainingRepo}
	commentHandler := &handlers.CommentHandler{Repository: commentRepo}
	favoriteHandler := &handlers.FavoriteHandler{Repository: favouriteRepo}

	// Регистрация HTTP обработчиков
	http.HandleFunc("/trainings", trainingHandler.GetAllTrainingsHandler)
	http.HandleFunc("/favorites/get", favoriteHandler.GetFavoriteHandler)
	http.HandleFunc("/favorites/add", favoriteHandler.AddToFavoritesHandler)
	http.HandleFunc("/favorites/remove", favoriteHandler.RemoveFromFavoritesHandler)
	http.HandleFunc("/comments", commentHandler.AddCommentHandler)
	http.HandleFunc("/comments/trainings", commentHandler.GetCommentsByTrainingIDHandler)
	http.HandleFunc("/rating/add", commentHandler.AddRatingHandler)
	http.HandleFunc("/rating/get", commentHandler.GetRatingHandler)
	http.HandleFunc("/trainings/rating", commentHandler.GetRatingHandler)

	// Запуск сервера
	http.ListenAndServe(":8085", nil)
}
