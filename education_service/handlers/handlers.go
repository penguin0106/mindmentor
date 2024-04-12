package handlers

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"mindmentor/education_service/models"
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

func GetEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Запрос к базе данных для получения информации об образовании
	rows, err := db.Query(context.Background(), "SELECT id, title, content FROM education")
	if err != nil {
		http.Error(w, "Failed to fetch education information", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var education []models.Education
	for rows.Next() {
		var edu models.Education
		err := rows.Scan(&edu.ID, &edu.Title, &edu.Content)
		if err != nil {
			http.Error(w, "Failed to fetch education information", http.StatusInternalServerError)
			return
		}
		education = append(education, edu)
	}

	json.NewEncoder(w).Encode(education)
}
