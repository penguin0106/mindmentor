package handlers

import (
	"encoding/json"
	"log"
	"meditation_service/services"
	"net/http"
	"strconv"
)

// RatingHandler handles HTTP requests related to ratings
type RatingHandler struct {
	RatingService *services.RatingService
}

func NewRatingHandler(ratService *services.RatingService) *RatingHandler {
	return &RatingHandler{
		RatingService: ratService,
	}
}

// AddRatingHandler adds a new rating for a course
func (h *RatingHandler) AddRatingHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		http.Error(w, "Некорректный идентификатор пользователя", http.StatusBadRequest)
		return
	}

	videoID, err := strconv.Atoi(r.FormValue("video_id"))
	if err != nil {
		http.Error(w, "Некорректный идентификатор видеофайла", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Некорректное значение оценки", http.StatusBadRequest)
		return
	}

	err = h.RatingService.AddVideoRating(userID, videoID, float64(value))
	if err != nil {
		http.Error(w, "Ошибка при добавлении оценки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *RatingHandler) GetAverageRatingHandler(w http.ResponseWriter, r *http.Request) {
	videoIDStr := r.URL.Query().Get("video_id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		log.Println("Error parsing video ID:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	averageRating, err := h.RatingService.GetVideoAverageRating(videoID)
	if err != nil {
		log.Println("Error getting average rating for video:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := struct {
		AverageRating float64 `json:"average_rating"`
	}{AverageRating: averageRating}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshalling average rating to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
