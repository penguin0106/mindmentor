package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"profile_service/services"
	"strconv"
)

type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler создает новый экземпляр UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUserHandler обрабатывает запрос на получение информации о пользователе по его ID
func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметра userID из запроса
	userIDStr := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Некорректный ID пользователя", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		log.Println("Error getting user by ID:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshalling user to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// EditProfileUsernameHandler обрабатывает запрос на изменение имени пользователя
func (h *UserHandler) EditProfileUsernameHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг данных запроса
	var requestData struct {
		UserID      int    `json:"user_id"`
		NewUsername string `json:"new_username"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Вызов сервиса для редактирования имени пользователя
	err = h.userService.EditProfileUsername(requestData.UserID, requestData.NewUsername)
	if err != nil {
		http.Error(w, "Ошибка при редактировании имени пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// EditProfileEmailHandler обрабатывает запрос на изменение email пользователя
func (h *UserHandler) EditProfileEmailHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг данных запроса
	var requestData struct {
		UserID   int    `json:"user_id"`
		NewEmail string `json:"new_email"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Вызов сервиса для редактирования email пользователя
	err = h.userService.EditProfileEmail(requestData.UserID, requestData.NewEmail)
	if err != nil {
		http.Error(w, "Ошибка при редактировании email пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// EditProfilePasswordHandler обрабатывает запрос на изменение пароля пользователя
func (h *UserHandler) EditProfilePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг данных запроса
	var requestData struct {
		UserID      int    `json:"user_id"`
		NewPassword string `json:"new_password"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Вызов сервиса для редактирования пароля пользователя
	err = h.userService.EditProfilePassword(requestData.UserID, requestData.NewPassword)
	if err != nil {
		http.Error(w, "Ошибка при редактировании пароля пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetFavoriteVideosHandler возвращает избранные видео пользователя
func (h *UserHandler) GetFavoriteVideosHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Некорректный ID пользователя", http.StatusBadRequest)
		return
	}

	favoriteVideos, err := h.userService.GetFavoriteVideos(userID)
	if err != nil {
		log.Println("Error getting favorite videos:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(favoriteVideos)
	if err != nil {
		log.Println("Error marshalling favorite videos to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetFavoriteTrainingsHandler возвращает избранные тренировки пользователя
func (h *UserHandler) GetFavoriteTrainingsHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Некорректный ID пользователя", http.StatusBadRequest)
		return
	}

	favoriteTrainings, err := h.userService.GetFavoriteTrainings(userID)
	if err != nil {
		log.Println("Error getting favorite trainings:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(favoriteTrainings)
	if err != nil {
		log.Println("Error marshalling favorite trainings to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
