package handlers

import (
	"fmt"
	"mindmentor/services/auth_service/repositories"
	"net/http"
)

// AuthorizationHandler обрабатывает запросы на авторизацию пользователей
func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	// Получить токен доступа из HTTP заголовка "Authorization"
	token := r.Header.Get("Authorization")

	// Создать экземпляр репозитория TokenRepository
	tokenRepo := &repositories.TokenRepository{DB: YourDBConnection}

	// Проверить валидность токена, используя метод репозитория
	isValid, err := tokenRepo.CheckTokenValidity(token)
	if err != nil {
		// Если произошла ошибка при проверке токена, вернуть статус 500 (Internal Server Error)
		http.Error(w, "Ошибка проверки токена", http.StatusInternalServerError)
		return
	}

	if !isValid {
		// Если токен не валиден, вернуть статус 401 (Unauthorized)
		http.Error(w, "Токен недействителен", http.StatusUnauthorized)
		return
	}

	// Если токен валиден, вернуть HTTP статус 200 (OK) и сообщение об успешной авторизации
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Пользователь успешно авторизован")
}
