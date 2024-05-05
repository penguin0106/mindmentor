package services

import (
	"meditation_service/models"
	"meditation_service/repositories"
)

type RatingService struct {
	RatingRepository *repositories.RatingRepository
}

func NewRatingsService(ratRepo *repositories.RatingRepository) *RatingService {
	return &RatingService{RatingRepository: ratRepo}
}

// AddCourseRating добавляет новую оценку курса медитации
func (s *RatingService) AddCourseRating(rating *models.Rating) error {
	return s.RatingRepository.AddRating(rating)
}

// GetAverageCourseRating возвращает среднюю оценку курса медитации
func (s *RatingService) GetAverageCourseRating(courseID int) (float64, error) {
	return s.RatingRepository.GetAverageRating(courseID)
}
