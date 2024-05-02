package middleware

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"mindmentor/services/auth_service/repositories"
	"mindmentor/shared/models"
	"net/http"
	"strconv"
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
	// Парсим JWT-токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем алгоритм подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неподдерживаемый алгоритм подписи")
		}
		// Возвращаем секретный ключ для проверки подписи
		return []byte("your_secret_key"), nil
	})
	if err != nil {
		return false
	}

	// Проверяем срок действия токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := time.Unix(int64(claims["exp"].(float64)), 0)
		if exp.Before(time.Now()) {
			return false
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return false
		}

		userRepo := &repositories.UserRepository{}

		exists, err := userRepo.CheckUserExists(userID)
		if err != nil || !exists {
			return false
		}
		return true

	}

	return false
}

// IsUserExistsByID Функция для получения пользователя из базы данных по ID
func IsUserExistsByID(userID int) bool {
	// Get user from the database
	user, err := GetUserByID(userID)
	if err != nil {
		// Handle database query error
		return false
	}
	if user == nil {
		// User not found in the database
		return false
	}
	// User found in the database
	return true
}

// GetUserByID retrieves a user from the database based on the provided ID
func GetUserByID(userID int) (*models.User, error) {
	// Connect to the database
	db, err := sql.Open("postgres", "your_database_connection_string")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// SQL query to get user information by ID
	query := "SELECT id, username, email FROM users WHERE id = $1"

	// Execute the SQL query
	var user models.User
	err = db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return nil, nil
		}
		// Error executing the query
		return nil, err
	}

	return &user, nil
}
