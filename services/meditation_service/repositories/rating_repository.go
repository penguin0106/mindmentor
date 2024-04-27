package repositories

import (
	"database/sql"
	"mindmentor/services/meditation_service/models"
)

// RatingRepository представляет собой репозиторий для работы с оценками курсов медитации
type RatingRepository struct {
	DB *sql.DB
}

// AddRating добавляет новую оценку курса медитации
func (r *RatingRepository) AddRating(rating *models.Rating) error {
	// Реализация добавления оценки в базу данных или другой источник данных
	_, err := r.DB.Exec("INSERT INTO ratings (course_id, user_id, value) VALUES ($1, $2, $3)", rating.CourseID, rating.UserID, rating.Value)
	return err
}

// GetAverageRating возвращает среднюю оценку курса медитации
func (r *RatingRepository) GetAverageRating(courseID int) (float64, error) {
	// Реализация запроса к базе данных или другому источнику данных для получения средней оценки курса
	var averageRating float64
	err := r.DB.QueryRow("SELECT AVG(value) FROM ratings WHERE course_id = $1", courseID).Scan(&averageRating)
	if err != nil {
		return 0, err
	}
	return averageRating, nil
}

// Other methods...
