package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"mindmentor/auth_service/models"
)

var db *pgx.Conn // Переменная для подключения к базе данных
var jwtKey = []byte("your_secret_key")

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Приводим логин к нижнему регистру
	user.Username = strings.ToLower(user.Username)

	// Проверка на уникальность логина и email
	var existingUser models.User
	err = db.QueryRow(context.Background(), "SELECT id FROM users WHERE lower(username) = $1 OR lower(email) = $2", user.Username, user.Email).Scan(&existingUser.ID)
	if err != pgx.ErrNoRows {
		http.Error(w, "User with such username/email already exists", http.StatusBadRequest)
		return
	}

	// Проверка пароля
	if user.Password != user.ConfirmPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	if !isPasswordValid(user.Password) {
		http.Error(w, "Password should be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character", http.StatusBadRequest)
		return
	}

	// Сохраняем пользователя в базе данных
	_, err = db.Exec(context.Background(), "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Unable to register user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func isPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}

	// Пароль должен содержать как минимум одну заглавную букву, одну строчную букву, одну цифру и один специальный символ
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	digitRegex := regexp.MustCompile(`[0-9]`)
	specialCharRegex := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)

	if !uppercaseRegex.MatchString(password) || !lowercaseRegex.MatchString(password) || !digitRegex.MatchString(password) || !specialCharRegex.MatchString(password) {
		return false
	}

	return true
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Приводим логин к нижнему регистру перед проверкой
	user.Username = strings.ToLower(user.Username)

	// Проверяем наличие логина и пароля в базе данных (без учета регистра)
	var dbUsername string
	err = db.QueryRow(context.Background(), "SELECT username FROM users WHERE lower(username) = $1 AND password = $2", user.Username, user.Password).Scan(&dbUsername)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Генерируем JWT токен
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to generate JWT token", http.StatusInternalServerError)
		return
	}

	// Отправляем токен в ответе
	response := map[string]string{"token": tokenString}
	json.NewEncoder(w).Encode(response)
}
