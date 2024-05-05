package handlers

import (
	"encoding/json"
	"net/http"
	"social_service/services"
)

type DiscussionHandler struct {
	Service *services.DiscussionService
}

func NewDiscussionHandler(service *services.DiscussionService) *DiscussionHandler {
	return &DiscussionHandler{Service: service}
}

// AddDiscussionHandler обработчик запроса на создание нового чата
func (h *DiscussionHandler) AddDiscussionHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение данных запроса (например, из JSON тела запроса)
	var requestData struct {
		Topic   string `json:"topic"`
		OwnerID int    `json:"ownerId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для создания чата
	discussion, err := h.Service.AddDiscussion(r.Context(), requestData.Topic, requestData.OwnerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа с данными созданного чата
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(discussion)
}

// FindDiscussionHandler обработчик запроса на поиск чата по теме
func (h *DiscussionHandler) FindDiscussionHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметра запроса (темы чата)
	topic := r.URL.Query().Get("topic")
	if topic == "" {
		http.Error(w, "Missing topic parameter", http.StatusBadRequest)
		return
	}

	// Вызов сервиса для поиска чата
	discussion, err := h.Service.FindDiscussionByTopic(r.Context(), topic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка найденного чата или сообщения об отсутствии чата с указанной темой
	w.Header().Set("Content-Type", "application/json")
	if discussion == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(discussion)
	}
}

// JoinDiscussionHandler обработчик запроса на присоединение пользователя к чату
func (h *DiscussionHandler) JoinDiscussionHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение данных запроса (например, из URL параметров или JSON тела запроса)
	var requestData struct {
		UserID       int `json:"userId"`
		DiscussionID int `json:"discussionId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для присоединения пользователя к чату
	if err := h.Service.JoinDiscussion(r.Context(), requestData.UserID, requestData.DiscussionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusOK)
}

// LeaveDiscussionHandler обработчик запроса на отсоединение пользователя от чата
func (h *DiscussionHandler) LeaveDiscussionHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение данных запроса (например, из URL параметров или JSON тела запроса)
	var requestData struct {
		UserID       int `json:"userId"`
		DiscussionID int `json:"discussionId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для отсоединения пользователя от чата
	if err := h.Service.LeaveDiscussion(r.Context(), requestData.UserID, requestData.DiscussionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusOK)
}
