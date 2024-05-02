package main

import (
	"mindmentor/services/meditation_service/handlers"
	"mindmentor/services/meditation_service/repositories"
	"net/http"
)

func main() {
	// Initialize repositories
	courseRepo := &repositories.CourseRepository{}
	musicRepo := &repositories.MusicRepository{}
	ratingRepo := &repositories.RatingRepository{}
	commentRepo := &repositories.CommentRepository{}
	favoriteRepo := &repositories.FavoriteRepository{}

	// Initialize handlers
	courseHandler := &handlers.CourseHandler{CourseRepo: courseRepo}
	musicHandler := &handlers.MusicHandler{MusicRepo: musicRepo}
	ratingHandler := &handlers.RatingHandler{RatingRepo: ratingRepo}
	commentHandler := &handlers.CommentHandler{CommentRepo: commentRepo}
	favoriteHandler := &handlers.FavouriteHandler{Repository: favoriteRepo}

	// Register HTTP handlers
	http.HandleFunc("/courses", courseHandler.GetAllCoursesHandler)
	http.HandleFunc("/music", musicHandler.GetAllMusicHandler)
	http.HandleFunc("/ratings/add", ratingHandler.AddRatingHandler)
	http.HandleFunc("/ratings/get", ratingHandler.GetAverageRatingHandler)
	http.HandleFunc("/comments", commentHandler.AddCommentHandler)
	http.HandleFunc("/comments/course", commentHandler.GetCommentsByCourseIDHandler)
	http.HandleFunc("favorites/get", favoriteHandler.GetFavoriteHandler)
	http.HandleFunc("/favorites/add", favoriteHandler.AddToFavouritesHandler)
	http.HandleFunc("/favorites/remove", favoriteHandler.RemoveFromFavouritesHandler)

	// Start the server
	http.ListenAndServe(":8083", nil)
}
