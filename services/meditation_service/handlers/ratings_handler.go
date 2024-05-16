package handlers

import (
	"encoding/json"
	"fmt"
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

// AddRatingHandler обрабатывает запрос на добавление рейтинга для видео
func (h *RatingHandler) AddRatingHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметров из запроса
	videoID, err := strconv.Atoi(r.URL.Query().Get("video_id"))
	if err != nil {
		http.Error(w, "Некорректный ID видео", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Некорректный ID пользователя", http.StatusBadRequest)
		return
	}

	rating, err := strconv.ParseFloat(r.URL.Query().Get("rating"), 64)
	if err != nil {
		http.Error(w, "Некорректный формат рейтинга", http.StatusBadRequest)
		return
	}

	// Вызов сервиса для добавления рейтинга
	err = h.RatingService.AddRating(videoID, userID, rating)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при добавлении рейтинга: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAverageRatingForVideoHandler обрабатывает запрос на получение среднего рейтинга для видео
func (h *RatingHandler) GetAverageRatingForVideoHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметров из запроса
	videoID, err := strconv.Atoi(r.URL.Query().Get("video_id"))
	if err != nil {
		http.Error(w, "Некорректный ID видео", http.StatusBadRequest)
		return
	}

	// Вызов сервиса для получения среднего рейтинга
	avgRating, err := h.RatingService.GetAverageRatingForVideo(videoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при получении среднего рейтинга: %v", err), http.StatusInternalServerError)
		return
	}

	// Формирование ответа
	response := struct {
		AverageRating float64 `json:"average_rating"`
	}{AverageRating: avgRating}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
