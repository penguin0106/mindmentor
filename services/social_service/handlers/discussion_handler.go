package handlers

import (
	"encoding/json"
	"fmt"
	"mindmentor/services/social_service/repositories"
	"mindmentor/shared/models"
	"net/http"
	"strconv"
)

type DiscussionHandler struct {
	Repo *repositories.DiscussionRepository
}

func (h *DiscussionHandler) CreateDiscussionHandler(w http.ResponseWriter, r *http.Request) {
	var discussion models.Discussion
	err := json.NewDecoder(r.Body).Decode(&discussion)
	if err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	err = h.Repo.CreateDiscussion(&discussion)
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

	discussion, err := h.Repo.FindDiscussion(topic)
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

	// Вызываем метод репозитория для присоединения пользователя к обсуждению
	err = h.Repo.JoinDiscussion(userIDInt, discussionIDInt)
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
	err = h.Repo.LeaveDiscussion(userIDInt, discussionIDInt)
	if err != nil {
		http.Error(w, "Ошибка при выходе пользователя из обсуждения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Пользователь успешно вышел из обсуждения")
}
