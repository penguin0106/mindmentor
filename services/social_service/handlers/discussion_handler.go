package handlers

import (
	"encoding/json"
	"fmt"
	"mindmentor/services/social_service/services"
	"mindmentor/shared/models"
	"net/http"
	"strconv"
)

type DiscussionHandler struct {
	SocialService *services.DiscussionService
}

func NewDiscussionHandler(discussionService *services.DiscussionService) *DiscussionHandler {
	return &DiscussionHandler{
		SocialService: discussionService,
	}
}

func (h *DiscussionHandler) CreateDiscussionHandler(w http.ResponseWriter, r *http.Request) {
	var discussion models.Discussion
	err := json.NewDecoder(r.Body).Decode(&discussion)
	if err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	err = h.SocialService.CreateDiscussion(&discussion)
	if err != nil {
		http.Error(w, "Ошибка при создании обсуждения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Обсуждение успешно создано")
}

func (h *DiscussionHandler) FindDiscussionHandler(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	if topic == "" {
		http.Error(w, "Не указана тема обсуждения", http.StatusBadRequest)
		return
	}

	discussion, err := h.SocialService.FindDiscussion(topic)
	if err != nil {
		http.Error(w, "Ошибка при поиске обсуждения", http.StatusInternalServerError)
		return
	}

	if discussion == nil {
		http.Error(w, "Обсуждение не найдено", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(discussion)
}

func (h *DiscussionHandler) JoinDiscussionHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя и ID обсуждения из параметров запроса
	userID := r.URL.Query().Get("userID")
	discussionID := r.URL.Query().Get("discussionID")

	// Проверяем, что оба ID указаны
	if userID == "" || discussionID == "" {
		http.Error(w, "Не указаны ID пользователя или обсуждения", http.StatusBadRequest)
		return
	}

	// Преобразуем ID в формат int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Некорректный ID пользователя", http.StatusBadRequest)
		return
	}

	discussionIDInt, err := strconv.Atoi(discussionID)
	if err != nil {
		http.Error(w, "Некорректный ID обсуждения", http.StatusBadRequest)
		return
	}

	// Получаем контекст из запроса
	ctx := r.Context()

	// Вызываем метод репозитория для присоединения пользователя к обсуждению
	err = h.SocialService.JoinDiscussion(ctx, userIDInt, discussionIDInt)
	if err != nil {
		http.Error(w, "Ошибка при присоединении пользователя к обсуждению", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Пользователь успешно присоединен к обсуждению")
}

func (h *DiscussionHandler) LeaveDiscussionHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя и ID обсуждения из параметров запроса
	userID := r.URL.Query().Get("userID")
	discussionID := r.URL.Query().Get("discussionID")

	// Проверяем, что оба ID указаны
	if userID == "" || discussionID == "" {
		http.Error(w, "Не указаны ID пользователя или обсуждения", http.StatusBadRequest)
		return
	}

	// Преобразуем ID в формат int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Некорректный ID пользователя", http.StatusBadRequest)
		return
	}

	discussionIDInt, err := strconv.Atoi(discussionID)
	if err != nil {
		http.Error(w, "Некорректный ID обсуждения", http.StatusBadRequest)
		return
	}

	// Вызываем метод репозитория для выхода пользователя из обсуждения
	err = h.SocialService.LeaveDiscussion(userIDInt, discussionIDInt)
	if err != nil {
		http.Error(w, "Ошибка при выходе пользователя из обсуждения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Пользователь успешно вышел из обсуждения")
}

func (h *DiscussionHandler) UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем данные сообщения из тела запроса
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Неверный формат данных сообщения", http.StatusBadRequest)
		return
	}

	// Получаем контекст из запроса
	ctx := r.Context()

	// Вызываем сервисный метод для обновления сообщения
	err = h.SocialService.UpdateMessage(ctx, &message)
	if err != nil {
		http.Error(w, "Ошибка при обновлении сообщения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Сообщение успешно обновлено")
}

func (h *DiscussionHandler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID сообщения из параметров запроса
	messageID := r.URL.Query().Get("messageID")

	// Проверяем, что ID сообщения указан
	if messageID == "" {
		http.Error(w, "Не указан ID сообщения", http.StatusBadRequest)
		return
	}

	// Преобразуем ID в формат int
	messageIDInt, err := strconv.Atoi(messageID)
	if err != nil {
		http.Error(w, "Некорректный ID сообщения", http.StatusBadRequest)
		return
	}

	// Получаем контекст из запроса
	ctx := r.Context()

	// Вызываем метод репозитория для удаления сообщения
	err = h.SocialService.DeleteMessage(ctx, messageIDInt)
	if err != nil {
		http.Error(w, "Ошибка при удалении сообщения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Сообщение успешно удалено")
}
