package handlers

import (
	"encoding/json"
	"fmt"
	"mindmentor/services/auth_service/models"
	"mindmentor/services/auth_service/services"
	"net/http"
)

// RegisterHandler обрабатывает запросы на регистрацию пользователей
func RegisterHandler(authService *services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Чтение данных запроса
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		// Регистрация пользователя
		if err := authService.RegisterUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Отправка успешного ответа
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User registered successfully")
	}
}

// LoginHandler обрабатывает запросы на аутентификацию пользователей
func LoginHandler(authService *services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Чтение данных запроса
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		// Аутентификация пользователя
		ok, err := authService.Authenticate(credentials.Email, credentials.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !ok {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Отправка успешного ответа
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User authenticated successfully")
	}
}
