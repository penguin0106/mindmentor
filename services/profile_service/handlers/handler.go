package handlers

import (
	"encoding/json"
	"net/http"
	"profile_service/models"
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
	userID := extractUserID(r)

	// Получение информации о пользователе из UserService
	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		// Обработка ошибки
		http.Error(w, "Failed to get user information", http.StatusInternalServerError)
		return
	}

	// Отправка информации о пользователе в формате JSON
	jsonResponse(w, user)
}

// UpdateUserHandler обрабатывает запрос на обновление информации о пользователе
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметра userID из запроса
	userID := extractUserID(r)

	// Извлечение данных пользователя из тела запроса
	updatedUser := extractUserFromRequest(r)

	// Обновление информации о пользователе с помощью UserService
	err := h.userService.UpdateUser(userID, updatedUser)
	if err != nil {
		// Обработка ошибки
		http.Error(w, "Failed to update user information", http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) GetFavoriteCourseHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Не указан идетификатор пользователя", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Некорректный формат идентификатора пользователя", http.StatusBadRequest)
		return
	}

	favorites, err := h.userService.GetFavoriteCourse(userID)
	if err != nil {
		http.Error(w, "Ошибка при получении избранных элементов пользователя", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(favorites)
	if err != nil {
		http.Error(w, "Ошибка при преобразовании данных в JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *UserHandler) GetFavoriteTrainingHandler(w http.ResponseWriter, r *http.Request) {
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

	favorites, err := h.userService.GetFavoriteTraining(userID)
	if err != nil {
		http.Error(w, "Ошибка при получении избранных элементов пользователя", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(favorites)
	if err != nil {
		http.Error(w, "Ошибка при преобразовании данных в JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func extractUserID(r *http.Request) int {
	// Пример извлечения параметра userID из URL запроса
	userIDParam := r.URL.Query().Get("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		// В случае ошибки возвращаем 0 или другое значение по умолчанию
		return 0
	}
	return userID
}

func extractUserFromRequest(r *http.Request) *models.User {
	// Пример извлечения данных пользователя из тела запроса
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Обработка ошибки, например, запись в лог или возвращение ошибки
		return nil
	}
	return &user
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	// Пример формирования ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		// Обработка ошибки, например, запись в лог или возвращение ошибки
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}
