package handlers

import (
	"fmt"
	"mindmentor/services/trainings_service/repositories"
	"net/http"
	"strconv"
)

// FavoriteHandler представляет обработчик HTTP-запросов для работы с избранными тренировками
type FavoriteHandler struct {
	Repository *repositories.FavoriteRepository
}

// AddToFavoritesHandler обрабатывает запрос на добавление тренировки в избранное
func (h *FavoriteHandler) AddToFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора пользователя из параметров запроса
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Не указан идентификатор пользователя", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Некорректный формат идентификатора пользователя", http.StatusBadRequest)
		return
	}

	// Получение идентификатора тренировки из параметров запроса
	trainingIDStr := r.URL.Query().Get("training_id")
	if trainingIDStr == "" {
		http.Error(w, "Не указан идентификатор тренировки", http.StatusBadRequest)
		return
	}

	trainingID, err := strconv.Atoi(trainingIDStr)
	if err != nil {
		http.Error(w, "Некорректный формат идентификатора тренировки", http.StatusBadRequest)
		return
	}

	// Добавление тренировки в избранное
	err = h.Repository.AddToFavorites(userID, trainingID)
	if err != nil {
		http.Error(w, "Ошибка добавления тренировки в избранное", http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Тренировка успешно добавлена в избранное")
}

// RemoveFromFavoritesHandler обрабатывает запрос на удаление тренировки из избранного
func (h *FavoriteHandler) RemoveFromFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора пользователя из параметров запроса
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Не указан идентификатор пользователя", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Некорректный формат идентификатора пользователя", http.StatusBadRequest)
		return
	}

	// Получение идентификатора тренировки из параметров запроса
	trainingIDStr := r.URL.Query().Get("training_id")
	if trainingIDStr == "" {
		http.Error(w, "Не указан идентификатор тренировки", http.StatusBadRequest)
		return
	}

	trainingID, err := strconv.Atoi(trainingIDStr)
	if err != nil {
		http.Error(w, "Некорректный формат идентификатора тренировки", http.StatusBadRequest)
		return
	}

	// Удаление тренировки из избранного
	err = h.Repository.RemoveFromFavorites(userID, trainingID)
	if err != nil {
		http.Error(w, "Ошибка удаления тренировки из избранного", http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Тренировка успешно удалена из избранного")
}
