package handlers

import (
	"encoding/json"
	"net/http"
	"social_service/models"
	"social_service/services"
)

type MessageHandler struct {
	Service *services.MessageService
}

func NewMessageHandler(service *services.MessageService) *MessageHandler {
	return &MessageHandler{Service: service}
}

// SendMessageHandler обработчик запроса на отправку сообщения
func (h *MessageHandler) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение данных запроса (например, из JSON тела запроса)
	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для отправки сообщения
	if err := h.Service.SendMessage(r.Context(), &message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusCreated)
}

// EditMessageHandler обработчик запроса на редактирование сообщения
func (h *MessageHandler) EditMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение данных запроса (например, из URL параметров или JSON тела запроса)
	var requestData struct {
		MessageID int    `json:"messageId"`
		Text      string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := &models.Message{
		ID:   requestData.MessageID,
		Text: requestData.Text,
	}

	// Вызов сервиса для редактирования сообщения
	if err := h.Service.EditMessage(r.Context(), message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusOK)
}

// DeleteMessageHandler обработчик запроса на удаление сообщения
func (h *MessageHandler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение данных запроса (например, из URL параметров или JSON тела запроса)
	var requestData struct {
		MessageID int `json:"messageId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для удаления сообщения
	if err := h.Service.DeleteMessage(r.Context(), requestData.MessageID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusOK)
}
