package handlers

import (
	"fmt"
	"mindmentor/services/meditation_service/repositories"
	"net/http"
	"strconv"
)

type FavouriteHandler struct {
	Repository *repositories.FavoriteRepository
}

func (h *FavouriteHandler) AddToFavouritesHandler(w http.ResponseWriter, r *http.Request) {

	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Не указан идентификатор пользователя", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Некорректный формат идетификатора пользователя", http.StatusBadRequest)
		return
	}

	courseIDStr := r.URL.Query().Get("course_id")
	if courseIDStr == "" {
		http.Error(w, "Не указан идетификатор курса", http.StatusBadRequest)
	}

	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		http.Error(w, "Некорректный формат идетификатора тренировки", http.StatusBadRequest)
		return
	}

	err = h.Repository.AddToFavorite(userID, courseID)
	if err != nil {
		http.Error(w, "Ошибка добавления тренировки в избранное", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Курс успешно добавлен в избранное")
}

func (h *FavouriteHandler) RemoveFromFavouritesHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Не указан идетификатор пользователя", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Некорректный формат идетификатора пользователя", http.StatusBadRequest)
		return
	}

	courseIDStr := r.URL.Query().Get("course_id")
	if courseIDStr == "" {
		http.Error(w, "Не указан идетификатор курса", http.StatusBadRequest)
		return
	}

	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		http.Error(w, "Некорректный формат идентификатора курса", http.StatusBadRequest)
		return
	}

	err = h.Repository.RemoveFromFavorite(userID, courseID)
	if err != nil {
		http.Error(w, "Ошибка удаления курса из избранного", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Курс успешно удален из избранного")
}
