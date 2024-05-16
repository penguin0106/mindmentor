package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"meditation_service/models"
	"time"
)

// RatingRepository представляет собой репозиторий для работы с оценками курсов медитации
type RatingRepository struct {
	DB *sql.DB
}

func NewRatingRepository(db *sql.DB) *RatingRepository {
	return &RatingRepository{DB: db}
}

// AddRating добавляет рейтинг для видео
func (r *RatingRepository) AddRating(rating *models.Rating) error {
	query := `INSERT INTO video_ratings (video_id, user_id, rating, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	var ratingID int
	err := r.DB.QueryRow(query, rating.VideoID, rating.UserID, rating.Rating, time.Now()).Scan(&ratingID)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении рейтинга: %v", err)
	}
	return nil
}

// GetAverageRatingForVideo возвращает средний рейтинг для указанного видео
func (r *RatingRepository) GetAverageRatingForVideo(videoID int) (float64, error) {
	query := `SELECT AVG(rating) FROM video_ratings WHERE video_id = $1`
	var avgRating float64
	err := r.DB.QueryRow(query, videoID).Scan(&avgRating)
	if err != nil {
		return 0, fmt.Errorf("ошибка при получении среднего рейтинга: %v", err)
	}
	return avgRating, nil
}

// GetUserRatingForVideo возвращает рейтинг пользователя для указанного видео
func (r *RatingRepository) GetUserRatingForVideo(videoID, userID int) (float64, error) {
	query := `SELECT rating FROM video_ratings WHERE video_id = $1 AND user_id = $2`
	var userRating float64
	err := r.DB.QueryRow(query, videoID, userID).Scan(&userRating)
	if err != nil {
		if err == sql.ErrNoRows {
			// Возвращаем ноль, если рейтинг не найден
			return 0, nil
		}
		return 0, fmt.Errorf("ошибка при получении рейтинга пользователя: %v", err)
	}
	return userRating, nil
}
