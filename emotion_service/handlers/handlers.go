package handlers

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"mindmentor/emotion_service/models"
	"net/http"
)

var db *pgx.Conn // Переменная для подключения к базе данных

func init() {
	// Инициализация соединения с базой данных при запуске приложения
	var err error
	db, err = pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/database")
	if err != nil {
		panic(err)
	}
}

func GetEmotions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Запрос к базе данных для получения списка эмоций
	rows, err := db.Query(context.Background(), "SELECT id, name FROM emotions")
	if err != nil {
		http.Error(w, "Failed to fetch emotions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var emotions []models.Emotion
	for rows.Next() {
		var emotion models.Emotion
		err := rows.Scan(&emotion.ID, &emotion.Name)
		if err != nil {
			http.Error(w, "Failed to fetch emotions", http.StatusInternalServerError)
			return
		}
		emotions = append(emotions, emotion)
	}

	json.NewEncoder(w).Encode(emotions)
}

func AddEmotion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emotion models.Emotion
	_ = json.NewDecoder(r.Body).Decode(&emotion)

	// Вставка эмоции в базу данных
	_, err := db.Exec(context.Background(), "INSERT INTO emotions (name) VALUES ($1)", emotion.Name)
	if err != nil {
		http.Error(w, "Failed to add emotion", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(emotion)
}
