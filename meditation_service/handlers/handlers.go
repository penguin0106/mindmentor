package handlers

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"mindmentor/meditation_service/models"
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

func GetMeditations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Запрос к базе данных для получения информации о практиках медитации
	rows, err := db.Query(context.Background(), "SELECT id, name, description, video_url, audio_url FROM meditations")
	if err != nil {
		http.Error(w, "Failed to fetch meditations", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var meditations []models.Meditation
	for rows.Next() {
		var meditation models.Meditation
		err := rows.Scan(&meditation.ID, &meditation.Name, &meditation.Description, &meditation.VideoURL, &meditation.AudioURL)
		if err != nil {
			http.Error(w, "Failed to fetch meditations", http.StatusInternalServerError)
			return
		}
		meditations = append(meditations, meditation)
	}

	json.NewEncoder(w).Encode(meditations)
}
