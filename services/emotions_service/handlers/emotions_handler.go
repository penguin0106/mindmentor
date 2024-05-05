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
	var emotion models.Emotion
	err := json.NewDecoder(r.Body).Decode(&emotion)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Дополнительные проверки и валидация данных эмоции

	err = h.EmotionService.CreateEmotion(&emotion)
	if err != nil {
		http.Error(w, "Ошибка создания записи эмоции", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emotion)
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
	emotionID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Некорректный ID эмоции", http.StatusBadRequest)
		return
	}

	err = h.EmotionService.DeleteEmotion(emotionID)
	if err != nil {
		http.Error(w, "Ошибка удаления записи эмоции", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetEmotionsByUserHandler обрабатывает запрос на получение эмоций пользователя
func (h *EmotionHandler) GetEmotionsByUserHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора пользователя из запроса
	userID, err := getUserIDFromRequest(r)
	if err != nil {
		// Обработка ошибки, если произошла
		http.Error(w, "Ошибка получения идентификатора пользователя", http.StatusBadRequest)
		return
	}

	// Получение эмоций пользователя из репозитория
	emotions, err := h.EmotionService.GetEmotionsByUserID(userID)
	if err != nil {
		http.Error(w, "Ошибка при получении эмоций пользователя", http.StatusInternalServerError)
		return
	}

	// Кодируем эмоции в JSON и отправляем клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emotions)
}

// getUserIDFromRequest извлекает идентификатор пользователя из запроса
func getUserIDFromRequest(r *http.Request) (int, error) {
	// Извлекаем параметр "user_id" из строки запроса
	userID := r.URL.Query().Get("user_id")

	// Проверяем, что параметр не пустой
	if userID == "" {
		return 0, nil // Возвращаем 0 в случае отсутствия или неверного идентификатора
	}

	// Преобразуем полученный идентификатор в целочисленное значение
	id, err := strconv.Atoi(userID)
	if err != nil {
		// Если произошла ошибка преобразования, возвращаем ошибку
		return 0, err
	}

	// Возвращаем полученный идентификатор пользователя
	return id, nil
}
