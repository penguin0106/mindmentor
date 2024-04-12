package handlers

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"mindmentor/activity_tracker_service/models"
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

func AddActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var activity models.Activity
	_ = json.NewDecoder(r.Body).Decode(&activity)

	// Вставка данных об активности в базу данных
	_, err := db.Exec(context.Background(), "INSERT INTO activity (user_id, steps, sleep_hours) VALUES ($1, $2, $3)", activity.UserID, activity.Steps, activity.SleepHours)
	if err != nil {
		http.Error(w, "Failed to add activity", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(activity)
}
