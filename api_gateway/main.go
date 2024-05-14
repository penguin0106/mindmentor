package main

import (
	"api_gateway/handlers"
	"api_gateway/middleware"
	"fmt"
	"log"
	"net/http"
)

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
	// Set up middleware
	authMiddleware := middleware.AuthMiddleware
	loggingMiddleware := middleware.LoggingMiddleware

	// Auth_service
	http.HandleFunc("/auth/register", loggingMiddleware(corsMiddleware(http.HandlerFunc(handlers.AuthRegisterHandler))))
	http.HandleFunc("/auth/login", loggingMiddleware(corsMiddleware(http.HandlerFunc(handlers.AuthLoginHandler))))

	//Emotions_service
	http.HandleFunc("/emotions/create", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.EmotionsCreateHandler)))))
	http.HandleFunc("/emotions/update", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.EmotionsUpdateHandler)))))
	http.HandleFunc("/emotions/delete", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.EmotionsDeleteHandler)))))
	http.HandleFunc("/emotions/user", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.EmotionsUserHandler)))))

	//Meditations_service
	http.HandleFunc("/meditations/video/all", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedVideoAllHandler)))))
	http.HandleFunc("/meditations/video/search", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedVideoSearchHandler)))))
	http.HandleFunc("/meditations/video/add", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedVideoAddHandler)))))
	http.HandleFunc("/meditations/video/delete", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedVideoDeleteHandler)))))

	http.HandleFunc("/meditations/music/all", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedMusicAllHandler)))))
	http.HandleFunc("/meditations/music/add", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedMusicAddHandler)))))

	http.HandleFunc("/meditations/ratings/get", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedRatingsGetHandler)))))
	http.HandleFunc("/meditations/ratings/add", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedRatingsAddHandler)))))

	http.HandleFunc("/meditations/comments/get", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedCommentsGetHandler)))))
	http.HandleFunc("/meditations/comments/add", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedCommentsAddHandler)))))

	http.HandleFunc("/meditations/favorites/add", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedFavouritesAddHandler)))))
	http.HandleFunc("/meditations/favorites/remove", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.MedFavouritesRemoveHandler)))))

	//Profile_service
	http.HandleFunc("/profile/user/get", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.ProfileUserGetHandler)))))
	http.HandleFunc("/profile/user/update", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.ProfileUserUpdateHandler)))))
	http.HandleFunc("/profile/favorites/course", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.ProfileFavouritesCoursesHandler)))))
	http.HandleFunc("/profile/favorites/training", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.ProfileFavouritesTrainingsHandler)))))

	//Social_service
	http.HandleFunc("/social/discussions/add", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.SocialDiscussionAddHandler)))))
	http.HandleFunc("/social/discussions/find", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.SocialDiscussionFindHandler)))))
	http.HandleFunc("/social/discussions/join", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.SocialDiscussionJoinHandler)))))
	http.HandleFunc("/social/discussions/leave", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.SocialDiscussionLeaveHandler)))))

	http.HandleFunc("/social/messages/send", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.SocialMessageSendHandler)))))
	http.HandleFunc("/social/messages/edit", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.SocialMessageEditHandler)))))
	http.HandleFunc("/social/messages/delete", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.SocialMessageDeleteHandler)))))

	//Trainings_service
	http.HandleFunc("/trainings/get", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.TrainingsGetHandler)))))
	http.HandleFunc("/trainings/search", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.TrainingsSearchHandler)))))

	http.HandleFunc("/trainings/favorites/add", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.TrainingsFavouritesAddHandler)))))
	http.HandleFunc("/trainings/favorites/remove", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.TrainingsFavouritesRemoveHandler)))))

	http.HandleFunc("/trainings/comments/add", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.TrainingsCommentsAddHandler)))))
	http.HandleFunc("/trainings/comments/get", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.TrainingsCommentsGetHandler)))))

	http.HandleFunc("/trainings/rating/add", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.TrainingsRatingAddHandler)))))
	http.HandleFunc("/trainings/rating/get", loggingMiddleware(authMiddleware(corsMiddleware(http.HandlerFunc(handlers.TrainingsRatingGetHandler)))))

	fmt.Println("API Gateway is running on port 8090...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
