package handlers

import (
	"encoding/json"
	"mindmentor/services/emotions_service/models"
	"mindmentor/services/emotions_service/repositories"
	"net/http"
	"strconv"
)

// EmotionHandler представляет обработчик HTTP-запросов для работы с эмоциями (записями)
type EmotionHandler struct {
	Repository *repositories.EmotionRepository
}

// CreateEmotionHandler обрабатывает запрос на создание новой эмоции (записи)
func (h *EmotionHandler) CreateEmotionHandler(w http.ResponseWriter, r *http.Request) {
	var emotion models.Emotion
	err := json.NewDecoder(r.Body).Decode(&emotion)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Дополнительные проверки и валидация данных эмоции

	err = h.Repository.CreateEmotion(&emotion)
	if err != nil {
		http.Error(w, "Ошибка создания эмоции", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emotion)
}

// GetEmotionsByUserHandler обрабатывает запрос на получение эмоций (записей) для указанного пользователя
func (h *EmotionHandler) GetEmotionsByUserHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора пользователя из запроса
	userID := getUserIDFromRequest(r)
	if userID == 0 {
		http.Error(w, "Не удалось получить идентификатор пользователя", http.StatusBadRequest)
		return
	}

	// Выполнение запроса к репозиторию для получения эмоций пользователя
	emotions, err := h.Repository.GetEmotionsByUserID(userID)
	if err != nil {
		http.Error(w, "Ошибка при получении эмоций пользователя", http.StatusInternalServerError)
		return
	}

	// Отправка эмоций в виде JSON-ответа
	responseJSON, err := json.Marshal(emotions)
	if err != nil {
		http.Error(w, "Ошибка при формировании JSON-ответа", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

// getUserIDFromRequest извлекает идентификатор пользователя из запроса
func getUserIDFromRequest(r *http.Request) int {
	// Извлекаем параметр "user_id" из строки запроса
	userID := r.URL.Query().Get("user_id")

	// Проверяем, что параметр не пустой
	if userID == "" {
		return 0 // Возвращаем 0 в случае отсутствия или неверного идентификатора
	}

	// Преобразуем полученный идентификатор в целочисленное значение
	id, err := strconv.Atoi(userID)
	if err != nil {
		// Если произошла ошибка преобразования, возвращаем 0
		return 0
	}

	// Возвращаем полученный идентификатор пользователя
	return id
}
