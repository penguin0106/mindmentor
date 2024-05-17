package handlers

import (
	"auth_service/models"
	"auth_service/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// AuthHandler представляет обработчики HTTP-запросов для аутентификации
type AuthHandler struct {
	AuthService *services.AuthService
	JWTService  *services.JWTService
}

// NewAuthHandler создает новый экземпляр обработчика аутентификации
func NewAuthHandler(authService *services.AuthService, jwtService *services.JWTService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		JWTService:  jwtService,
	}
}

// RegisterUserHandler обрабатывает запрос на регистрацию нового пользователя
func (handler *AuthHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRegistrationRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&userRegistrationRequest)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	err = handler.AuthService.RegisterUser(userRegistrationRequest.Username, userRegistrationRequest.Email, userRegistrationRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// AuthenticateUserHandler обрабатывает запрос на аутентификацию пользователя и генерацию JWT токена
func (handler *AuthHandler) AuthenticateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userAuthenticationRequest struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&userAuthenticationRequest)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Аутентификация пользователя
	user, err := handler.AuthService.AuthenticateUser(userAuthenticationRequest.Identifier, userAuthenticationRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Генерация JWT токена
	token, err := handler.JWTService.GenerateToken(user.ID, user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Отправка токена в ответе
	response := models.Token{Token: token}
	json.NewEncoder(w).Encode(response)
	fmt.Fprintf(w, "Пользвователь успешно авторизирован")
}

// VerifyTokenHandler обрабатывает запрос на верификацию JWT токена
func (handler *AuthHandler) VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("Не удалось разобрать тело запроса:", err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	userID, username, _, err := handler.JWTService.VerifyToken(requestBody.Token)
	if err != nil {
		log.Println("Ошибка проверки токена:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	response := struct {
		UserID   int    `json:"userID"`
		Username string `json:"username"`
	}{
		UserID:   userID,
		Username: username,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Не удалось закодировать ответ:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
