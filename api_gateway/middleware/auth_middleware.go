package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const authServiceURL = "http://auth-service:8081"

// AuthMiddleware выполняет проверку аутентификации пользователя
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		reqBody, err := json.Marshal(map[string]string{"token": token})
		if err != nil {
			log.Println("Ошибка при маршаллинге JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp, err := http.Post(authServiceURL+"/verify-token", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Println("Ошибка при отправке запроса в auth service:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("Статус ответа auth service:", resp.StatusCode)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
