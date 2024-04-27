package handlers

import (
	"encoding/json"
	"log"
	"mindmentor/services/meditation_service/repositories"
	"net/http"
)

// MusicHandler handles HTTP requests related to meditation music
type MusicHandler struct {
	MusicRepo *repositories.MusicRepository
}

// GetAllMusicHandler returns all meditation music
func (h *MusicHandler) GetAllMusicHandler(w http.ResponseWriter, r *http.Request) {
	music, err := h.MusicRepo.GetAllMusic()
	if err != nil {
		log.Println("Error getting all music:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(music)
	if err != nil {
		log.Println("Error marshalling music to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
