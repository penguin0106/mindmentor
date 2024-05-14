package repositories

import (
	"database/sql"
	"meditation_service/models"
)

// RatingRepository представляет собой репозиторий для работы с оценками курсов медитации
type RatingRepository struct {
	DB *sql.DB
}

func NewRatingrepository(db *sql.DB) *RatingRepository {
	return &RatingRepository{DB: db}
}

// AddRating добавляет новую оценку курса медитации
func (r *RatingRepository) AddRating(rating *models.Rating) error {
	// Реализация добавления оценки в базу данных или другой источник данных
	_, err := r.DB.Exec("INSERT INTO video_ratings (video_id, user_id, value) VALUES ($1, $2, $3)", rating.ItemID, rating.UserID, rating.Value)
	return err
}

// GetAverageRating возвращает среднюю оценку курса медитации
func (r *RatingRepository) GetAverageRating(videoID int) (float64, error) {
	// Реализация запроса к базе данных или другому источнику данных для получения средней оценки курса
	var averageRating float64
	err := r.DB.QueryRow("SELECT AVG(value) FROM video_ratings WHERE video_id = $1", videoID).Scan(&averageRating)
	if err != nil {
		return 0, err
	}
	return averageRating, nil
}
