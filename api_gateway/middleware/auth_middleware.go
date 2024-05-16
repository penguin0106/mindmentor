package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const authServiceURL = "http://auth-service:8081"

// AuthMiddleware выполняет проверку аутентификации пользователя
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получение токена из заголовка Authorization
		token := r.Header.Get("Authorization")

		// Проверка наличия токена
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Проверка валидности токена через auth-service
		reqBody, err := json.Marshal(map[string]string{"token": token})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp, err := http.Post(authServiceURL+"/verify-token", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Продолжение выполнения цепочки обработчиков
		next.ServeHTTP(w, r)
	})
}
