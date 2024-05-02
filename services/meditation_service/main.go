package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"mindmentor/services/meditation_service/handlers"
	"mindmentor/services/meditation_service/repositories"
	"mindmentor/services/meditation_service/services"
	"net/http"
)

func main() {
	// Подключение к базе данных
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Initialize repositories
	courseRepo := repositories.NewCourseRepository(db)
	musicRepo := repositories.NewMusicrepository(db)
	ratingRepo := repositories.NewRatingrepository(db)
	commentRepo := repositories.NewCommentrepository(db)
	favoriteRepo := repositories.NewFavoriterepository(db)

	courseServ := services.NewCourseService(courseRepo)
	musicServ := services.NewMusicService(musicRepo)
	ratingServ := services.NewRatingsService(ratingRepo)
	commentServ := services.NewCommentService(commentRepo)
	favoriteServ := services.NewFavoriteService(favoriteRepo)

	// Initialize handlers
	courseHandler := handlers.NewCourseHandler(courseServ)
	musicHandler := handlers.NewMusicHandler(musicServ)
	ratingHandler := handlers.NewRatingHandler(ratingServ)
	commentHandler := handlers.NewCommentHandler(commentServ)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteServ)

	// Register HTTP handlers
	http.HandleFunc("/courses", courseHandler.GetAllCoursesHandler)
	http.HandleFunc("/course/search", courseHandler.GetCourseByNameHandler)
	http.HandleFunc("course/add", courseHandler.AddCourseHandler)
	http.HandleFunc("/music", musicHandler.GetAllMusicHandler)
	http.HandleFunc("/music/add", musicHandler.AddMusicHandler)
	http.HandleFunc("/ratings/add", ratingHandler.AddRatingHandler)
	http.HandleFunc("/ratings/get", ratingHandler.GetAverageRatingHandler)
	http.HandleFunc("/comments", commentHandler.AddCommentHandler)
	http.HandleFunc("/comments/course", commentHandler.GetCommentsByCourseIDHandler)
	http.HandleFunc("/favorites/add", favoriteHandler.AddToFavouritesHandler)
	http.HandleFunc("/favorites/remove", favoriteHandler.RemoveFromFavouritesHandler)

	// Start the server
	http.ListenAndServe(":8083", nil)
}

// connectToDatabase подключается к базе данных и возвращает объект подключения
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@localhost:5432/mindmentor?sslmode=disable")
	return db, err
}
