package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"mindmentor/database/models"
	"net/http"
)

// DB представляет собой объект базы данных
var DB *sql.DB

// InitDB инициализирует объект базы данных
func InitDB(db *sql.DB) {
	DB = db
}

// GetUsers возвращает список всех пользователей
func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		log.Println("Error fetching users: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			log.Println("Error scanning user row: ", err)
			continue
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over user rows: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// CreateUser создает нового пользователя
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error decoding request body: ", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Проверяем, что все поля заполнены
	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Username, email, and password are required fields", http.StatusBadRequest)
		return
	}

	// Вставляем нового пользователя в базу данных
	_, err := DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		log.Println("Error inserting user into database: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateUser обновляет информацию о пользователе
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error decoding request body: ", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Обновляем информацию о пользователе в базе данных
	_, err := DB.Exec("UPDATE users SET username=$1, email=$2, password=$3 WHERE id=$4", user.Username, user.Email, user.Password, userID)
	if err != nil {
		log.Println("Error updating user in database: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUser удаляет пользователя
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	// Удаляем пользователя из базы данных
	_, err := DB.Exec("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		log.Println("Error deleting user from database: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetUser возвращает информацию о конкретном пользователе
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User
	err := DB.QueryRow("SELECT id, username, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		log.Println("Error fetching user: ", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
