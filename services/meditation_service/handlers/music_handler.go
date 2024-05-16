package handlers

import (
	"encoding/json"
	"log"
	"meditation_service/services"
	"net/http"
)

// MusicHandler handles HTTP requests related to meditation music
type MusicHandler struct {
	MusicServ *services.MusicService
}

func NewMusicHandler(musicServ *services.MusicService) *MusicHandler {
	return &MusicHandler{MusicServ: musicServ}
}

// GetAllMusicHandler возвращает всю музыку для медитации
func (h *MusicHandler) GetAllMusicHandler(w http.ResponseWriter, _ *http.Request) {
	musicList, err := h.MusicServ.GetAllMusic()
	if err != nil {
		log.Println("Ошибка при получении всей музыки:", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(musicList)
	if err != nil {
		log.Println("Ошибка при преобразовании музыки в JSON:", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// AddMusicHandler добавляет новый аудиофайл для медитации
func (h *MusicHandler) AddMusicHandler(w http.ResponseWriter, r *http.Request) {
	var musicData struct {
		Name     string `json:"name"`
		Duration int    `json:"duration"`
		Music    []byte `json:"music"`
	}

	err := json.NewDecoder(r.Body).Decode(&musicData)
	if err != nil {
		http.Error(w, "Ошибка декодирования тела запроса", http.StatusBadRequest)
		return
	}

	err = h.MusicServ.AddMusic(musicData.Name, musicData.Duration, musicData.Music)
	if err != nil {
		http.Error(w, "Ошибка при добавлении музыки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
