package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"meditation_service/handlers"
	"meditation_service/repositories"
	"meditation_service/services"
	"net/http"
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
	// Подключение к базе данных
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Initialize repositories
	videoRepo := repositories.NewVideoRepository(db)
	musicRepo := repositories.NewMusicRepository(db)
	ratingRepo := repositories.NewRatingRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	favoriteRepo := repositories.NewFavoriteRepository(db)

	videoServ := services.NewVideoService(videoRepo)
	musicServ := services.NewMusicService(musicRepo)
	ratingServ := services.NewRatingsService(ratingRepo)
	commentServ := services.NewCommentService(commentRepo)
	favoriteServ := services.NewFavoriteService(favoriteRepo)

	// Initialize handlers
	videoHandler := handlers.NewVideoHandler(videoServ)
	musicHandler := handlers.NewMusicHandler(musicServ)
	ratingHandler := handlers.NewRatingHandler(ratingServ)
	commentHandler := handlers.NewCommentHandler(commentServ)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteServ)

	// Register HTTP handlers
	http.Handle("/video/all", corsMiddleware(http.HandlerFunc(videoHandler.GetAllVideosHandler)))
	http.Handle("/video/search", corsMiddleware(http.HandlerFunc(videoHandler.GetVideoByTitleHandler)))
	http.Handle("/video/add", corsMiddleware(http.HandlerFunc(videoHandler.AddVideoHandler)))
	http.Handle("/video/delete", corsMiddleware(http.HandlerFunc(videoHandler.DeleteVideoHandler)))

	http.Handle("/music/all", corsMiddleware(http.HandlerFunc(musicHandler.GetAllMusicHandler)))
	http.Handle("/music/add", corsMiddleware(http.HandlerFunc(musicHandler.AddMusicHandler)))

	http.Handle("/ratings/add", corsMiddleware(http.HandlerFunc(ratingHandler.AddRatingHandler)))
	http.Handle("/ratings/get", corsMiddleware(http.HandlerFunc(ratingHandler.GetAverageRatingForVideoHandler)))

	http.Handle("/comments/add", corsMiddleware(http.HandlerFunc(commentHandler.AddCommentHandler)))
	http.Handle("/comments/get", corsMiddleware(http.HandlerFunc(commentHandler.GetCommentsByVideoIDHandler)))

	http.Handle("/favorites/add", corsMiddleware(http.HandlerFunc(favoriteHandler.AddToFavouriteHandler)))
	http.Handle("/favorites/remove", corsMiddleware(http.HandlerFunc(favoriteHandler.RemoveFromFavoriteHandler)))

	// Start the server
	http.ListenAndServe(":8083", nil)
}
