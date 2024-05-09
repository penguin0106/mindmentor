package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"trainings_service/handlers"
	"trainings_service/repositories"
	"trainings_service/services"
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
	http.Handle("/trainings", corsMiddleware(http.HandlerFunc(trainingHandler.GetAllTrainingsHandler)))
	http.Handle("/trainings/search", corsMiddleware(http.HandlerFunc(trainingHandler.GetTrainingByNameHandler)))

	// Register favorite handler functions with CORS middleware
	http.Handle("/favorites/add", corsMiddleware(http.HandlerFunc(favoriteHandler.AddToFavoritesHandler)))
	http.Handle("/favorites/remove", corsMiddleware(http.HandlerFunc(favoriteHandler.RemoveFromFavoritesHandler)))

	// Register comment handler functions with CORS middleware
	http.Handle("/comments", corsMiddleware(http.HandlerFunc(commentHandler.AddCommentHandler)))
	http.Handle("/comments/trainings", corsMiddleware(http.HandlerFunc(commentHandler.GetCommentsByTrainingIDHandler)))
	http.Handle("/rating/add", corsMiddleware(http.HandlerFunc(commentHandler.AddRatingHandler)))
	http.Handle("/rating/get", corsMiddleware(http.HandlerFunc(commentHandler.GetRatingHandler)))

	// Запуск сервера
	http.ListenAndServe(":8085", nil)
}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@localhost:5432/mindmentor?sslmode=disable")
	return db, err
}
