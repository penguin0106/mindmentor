package handlers

import (
	"encoding/json"
	"log"
	"mindmentor/services/meditation_service/models"
	"mindmentor/services/meditation_service/services"
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
	var rating models.Rating
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		log.Println("Error decoding rating JSON:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = h.RatingService.AddCourseRating(&rating)
	if err != nil {
		log.Println("Error adding rating:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *RatingHandler) GetAverageRatingHandler(w http.ResponseWriter, r *http.Request) {
	courseID, err := strconv.Atoi(r.URL.Query().Get("course_id"))
	if err != nil {
		http.Error(w, "Некорректный идентификатор тренировки", http.StatusBadRequest)
		return
	}
	averageRating, err := h.RatingService.GetAverageCourseRating(courseID)
	if err != nil {
		http.Error(w, "Ошибка при получении рейтинга тренировки", http.StatusInternalServerError)
		return
	}

	response := map[string]float64{"average_rating": averageRating}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Ошибка при кодировке ответа:", err)
	}
}
