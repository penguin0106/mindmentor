package handlers

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"mindmentor/trainings_service/models"
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

func SearchTrainings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// TODO: Implement searching trainings
}

func AddReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var review models.Review
	_ = json.NewDecoder(r.Body).Decode(&review)

	// Вставка нового отзыва о тренировке в базу данных
	_, err := db.Exec(context.Background(), "INSERT INTO reviews (training_id, user_id, rating, description) VALUES ($1, $2, $3, $4)", review.TrainingID, review.UserID, review.Rating, review.Description)
	if err != nil {
		http.Error(w, "Failed to add review", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(review)
}
