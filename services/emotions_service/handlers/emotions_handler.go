package handlers

import (
	"emotions_service/models"
	"emotions_service/services"
	"encoding/json"
	"net/http"
	"strconv"
)

// EmotionHandler представляет обработчик HTTP-запросов для работы с эмоциями (записями)
type EmotionHandler struct {
	EmotionService *services.EmotionService
}

func NewEmotionHandler(emoService *services.EmotionService) *EmotionHandler {
	return &EmotionHandler{
		EmotionService: emoService,
	}
}

// CreateEmotionHandler обрабатывает запрос на создание новой записи эмоции
func (h *EmotionHandler) CreateEmotionHandler(w http.ResponseWriter, r *http.Request) {
	var emotionRequest struct {
		Topic string `json:"topic"`
		Body  string `json:"body"`
	}
	err := json.NewDecoder(r.Body).Decode(&emotionRequest)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err = h.EmotionService.CreateEmotion(emotionRequest.Topic, emotionRequest.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateEmotionHandler обрабатывает запрос на обновление записи эмоции
func (h *EmotionHandler) UpdateEmotionHandler(w http.ResponseWriter, r *http.Request) {
	emotionID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Некорректный ID эмоции", http.StatusBadRequest)
		return
	}

	var updatedEmotion models.Emotion
	err = json.NewDecoder(r.Body).Decode(&updatedEmotion)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Дополнительные проверки и валидация данных обновленной эмоции

	err = h.EmotionService.UpdateEmotion(emotionID, &updatedEmotion)
	if err != nil {
		http.Error(w, "Ошибка обновления записи эмоции", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteEmotionHandler обрабатывает запрос на удаление записи эмоции
func (h *EmotionHandler) DeleteEmotionHandler(w http.ResponseWriter, r *http.Request) {
	emotionIDStr := r.URL.Query().Get("id")
	emotionID, err := strconv.Atoi(emotionIDStr)
	if err != nil {
		// В случае некорректного идентификатора, возвращаем ошибку и код 400 (BadRequest)
		http.Error(w, "Некорректный ID эмоции", http.StatusBadRequest)
		return
	}

	// Вызываем метод DeleteEmotion сервиса EmotionService для удаления эмоции
	err = h.EmotionService.DeleteEmotion(emotionID)
	if err != nil {
		// В случае ошибки при удалении эмоции, возвращаем ошибку и код 500 (InternalServerError)
		http.Error(w, "Ошибка удаления записи эмоции", http.StatusInternalServerError)
		return
	}

	// В случае успешного удаления, возвращаем код 200 (OK)
	w.WriteHeader(http.StatusOK)
}

// GetEmotionsByUserHandler обрабатывает запрос на получение эмоций пользователя
func (h *EmotionHandler) GetEmotionsByUserHandler(w http.ResponseWriter, _ *http.Request) {

	// Получение эмоций пользователя из репозитория
	emotions, err := h.EmotionService.GetEmotionsByUserID()
	if err != nil {
		http.Error(w, "Ошибка при получении эмоций пользователя", http.StatusInternalServerError)
		return
	}

	// Кодируем эмоции в JSON и отправляем клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emotions)
}
