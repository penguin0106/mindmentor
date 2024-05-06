package middleware

import (
	"api_gateway/models"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// AuthMiddleware выполняет проверку аутентификации пользователя
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверка аутентификации пользователя
		if !IsAuthenticated(r) {
			// Если пользователь не аутентифицирован, отправляем ошибку или перенаправляем на страницу входа
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Пользователь аутентифицирован, продолжаем выполнение цепочки обработчиков
		next.ServeHTTP(w, r)
	})
}

// IsAuthenticated проверяет аутентификацию пользователя
func IsAuthenticated(r *http.Request) bool {
	// Получаем токен из заголовка Authorization
	token := r.Header.Get("Authorization")

	// Проверяем, что токен не пустой
	if token == "" {
		return false
	}
	// Проверяем валидность токена
	if !IsTokenValid(token) {
		return false
	}

	// Проверяем существование пользователя в базе данных
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil || !IsUserExistsByID(userID) {
		return false
	}
	// Возвращаем true, если токен валиден
	return true
}

// IsTokenValid Функция для проверки валидности токена
func IsTokenValid(tokenString string) bool {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return false
	}

	// Декодируем заголовок
	_, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}

	// Декодируем полезную нагрузку
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	// Проверяем срок действия токена
	var claims models.TokenClaims
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return false
	}
	expirationTime := time.Unix(claims.Exp, 0)
	if time.Now().After(expirationTime) {
		return false
	}

	// Проверяем подпись токена
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return false
	}
	expectedSignature := computeHMAC(parts[0]+"."+parts[1], []byte("your_secret_key"))
	if !hmac.Equal(expectedSignature, signature) {
		return false
	}

	// Проверяем существование пользователя в базе данных
	userID, err := strconv.Atoi(claims.UserID)
	if err != nil || !IsUserExistsByID(userID) {
		return false
	}

	return true
}

// Функция для вычисления HMAC для подписи токена
func computeHMAC(data string, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return h.Sum(nil)
}

// IsUserExistsByID Функция для проверки существования пользователя в базе данных по ID
func IsUserExistsByID(userID int) bool {
	// Получаем пользователя из базы данных
	user, err := GetUserByID(userID)
	if err != nil {
		// Обработка ошибки запроса к базе данных
		return false
	}
	if user == nil {
		// Пользователь не найден в базе данных
		return false
	}
	// Пользователь найден в базе данных
	return true
}

// GetUserByID Функция для получения пользователя из базы данных по ID
func GetUserByID(userID int) (*models.User, error) {
	// Подключение к базе данных
	db, err := sql.Open("postgres", "postgres://mindmentor:postgres@database_service:5432/mindmentor?sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// SQL-запрос для получения информации о пользователе по ID
	query := "SELECT id, username, email FROM users WHERE id = $1"

	// Выполнение SQL-запроса
	var user models.User
	err = db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// Пользователь не найден
			return nil, nil
		}
		// Ошибка при выполнении запроса
		return nil, err
	}

	return &user, nil
}
