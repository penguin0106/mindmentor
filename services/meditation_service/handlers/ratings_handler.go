package handlers

import (
	"encoding/json"
	"log"
	"mindmentor/services/meditation_service/repositories"
	"mindmentor/shared/models"
	"net/http"
)

// RatingHandler handles HTTP requests related to ratings
type RatingHandler struct {
	RatingRepo *repositories.RatingRepository
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

	err = h.RatingRepo.AddRating(&rating)
	if err != nil {
		log.Println("Error adding rating:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
