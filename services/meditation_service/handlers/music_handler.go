package handlers

import (
	"encoding/json"
	"log"
	"mindmentor/services/meditation_service/services"
	"mindmentor/shared/models"
	"net/http"
)

// MusicHandler handles HTTP requests related to meditation music
type MusicHandler struct {
	MusicServ *services.MusicService
}

func NewMusicHandler(musicServ *services.MusicService) *MusicHandler {
	return &MusicHandler{MusicServ: musicServ}
}

// GetAllMusicHandler returns all meditation music
func (h *MusicHandler) GetAllMusicHandler(w http.ResponseWriter, _ *http.Request) {
	music, err := h.MusicServ.GetAllMusic()
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

// AddMusicHandler добавляет новый аудиофайл для медитации
func (h *MusicHandler) AddMusicHandler(w http.ResponseWriter, r *http.Request) {
	var music models.Music
	err := json.NewDecoder(r.Body).Decode(&music)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = h.MusicServ.AddMusic(&music)
	if err != nil {
		http.Error(w, "Failed to add music", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
