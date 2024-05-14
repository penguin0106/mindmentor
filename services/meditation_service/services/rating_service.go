package services

import (
	"errors"
	"meditation_service/models"
	"meditation_service/repositories"
)

type RatingService struct {
	RatingRepository *repositories.RatingRepository
}

func NewRatingsService(ratRepo *repositories.RatingRepository) *RatingService {
	return &RatingService{RatingRepository: ratRepo}
}

// AddVideoRating добавляет новую оценку курса медитации
func (s *RatingService) AddVideoRating(userID, videoID int, value float64) error {
	if value < 0 || value > 5 {
		return errors.New("недопустимое значение оценки")
	}
	rating := &models.Rating{
		UserID: userID,
		ItemID: videoID,
		Value:  value,
	}

	err := s.RatingRepository.AddRating(rating)
	if err != nil {
		return err
	}
	return nil
}

// GetVideoAverageRating возвращает среднюю оценку курса медитации
func (s *RatingService) GetVideoAverageRating(videoID int) (float64, error) {
	averageRating, err := s.RatingRepository.GetAverageRating(videoID)
	if err != nil {
		return 0, err
	}
	return averageRating, nil
}
