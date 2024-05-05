package handlers

import (
	"auth_service/services"
	"encoding/json"
	"net/http"
)

// AuthHandler представляет обработчики HTTP-запросов для аутентификации
type AuthHandler struct {
	AuthService *services.AuthService
}

// NewAuthHandler создает новый экземпляр обработчика аутентификации
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
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

// AuthenticateUserHandler обрабатывает запрос на аутентификацию пользователя
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

	user, err := handler.AuthService.AuthenticateUser(userAuthenticationRequest.Identifier, userAuthenticationRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(user)
}
