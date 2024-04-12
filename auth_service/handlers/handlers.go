package handlers

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"mindmentor/auth_service/models"
	"net/http"
)

var db *pgx.Conn // Переменная для подключения к базе данных

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/database")
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	// TODO: Валидация данных пользователя

	// Сохраняем пользователя в базе данных
	_, err = conn.Exec(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		http.Error(w, "Unable to register user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/database")
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	// Проверяем учетные данные пользователя в базе данных
	err = conn.QueryRow(context.Background(), "SELECT username, password FROM users WHERE username = $1 AND password = $2", user.Username, user.Password).Scan(&user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Генерируем JWT токен и отправляем его в ответе
	json.NewEncoder(w).Encode(map[string]string{"token": "your_generated_jwt_token_here"})
}
