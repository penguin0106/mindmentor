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

	// Initialize handlers
	courseHandler := &handlers.CourseHandler{CourseRepo: courseRepo}
	musicHandler := &handlers.MusicHandler{MusicRepo: musicRepo}
	ratingHandler := &handlers.RatingHandler{RatingRepo: ratingRepo}
	commentHandler := &handlers.CommentHandler{CommentRepo: commentRepo}

	// Register HTTP handlers
	http.HandleFunc("/courses", courseHandler.GetAllCoursesHandler)
	http.HandleFunc("/music", musicHandler.GetAllMusicHandler)
	http.HandleFunc("/ratings", ratingHandler.AddRatingHandler)
	http.HandleFunc("/comments", commentHandler.AddCommentHandler)
	http.HandleFunc("/comments/course", commentHandler.GetCommentsByCourseIDHandler)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
