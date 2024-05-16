package handlers

import (
	"encoding/json"
	"meditation_service/services"
	"net/http"
	"strconv"
)

type VideoHandler struct {
	VideoService *services.VideoService
}

func NewVideoHandler(videoService *services.VideoService) *VideoHandler {
	return &VideoHandler{
		VideoService: videoService,
	}
}

func (h *VideoHandler) GetAllVideosHandler(w http.ResponseWriter, _ *http.Request) {
	videos, err := h.VideoService.GetAllVideos()
	if err != nil {
		http.Error(w, "Ошибка при получении видеофайла", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}

func (h *VideoHandler) AddVideoHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		VideoContent []byte `json:"video_content"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err = h.VideoService.AddVideo(requestData.Title, requestData.Description, requestData.VideoContent)
	if err != nil {
		http.Error(w, "Ошибка при добавлении видеофайла", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *VideoHandler) GetVideoByIDHandler(w http.ResponseWriter, r *http.Request) {
	videoID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Некорректный ID видеофайла", http.StatusBadRequest)
		return
	}

	video, err := h.VideoService.GetVideoByID(videoID)
	if err != nil {
		http.Error(w, "Ошибка при получении видеофайла", http.StatusInternalServerError)
		return
	}

	if video == nil {
		http.Error(w, "Видеофайл с указанным идентификатором не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(video)
}

func (h *VideoHandler) DeleteVideoHandler(w http.ResponseWriter, r *http.Request) {
	videoID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Некорректный ID видеофайла", http.StatusBadRequest)
		return
	}

	err = h.VideoService.DeleteVideo(videoID)
	if err != nil {
		http.Error(w, "Ошибка при удалении видефайла", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *VideoHandler) GetVideoByTitleHandler(w http.ResponseWriter, r *http.Request) {
	videoTitle := r.URL.Query().Get("title")
	if videoTitle == "" {
		http.Error(w, "Не указано название видеофайла", http.StatusBadRequest)
		return
	}

	video, err := h.VideoService.GetVideoByTitle(videoTitle)
	if err != nil {
		http.Error(w, "Ошибка при получении видеофайла", http.StatusInternalServerError)
		return
	}

	if video == nil {
		http.Error(w, "Видеофайл с указанным названием не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(video)
}
