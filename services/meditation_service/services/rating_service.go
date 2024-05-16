package services

import (
	"fmt"
	"meditation_service/models"
	"meditation_service/repositories"
)

type RatingService struct {
	RatingRepository *repositories.RatingRepository
}

func NewRatingsService(ratRepo *repositories.RatingRepository) *RatingService {
	return &RatingService{RatingRepository: ratRepo}
}

// AddRating добавляет рейтинг для указанного видео и пользователя
func (s *RatingService) AddRating(videoID, userID int, rating float64) error {
	// Проверяем, существует ли уже рейтинг для данного пользователя и видео
	userRating, err := s.RatingRepository.GetUserRatingForVideo(videoID, userID)
	if err != nil {
		return fmt.Errorf("ошибка при проверке рейтинга пользователя: %v", err)
	}

	if userRating != 0 {
		return fmt.Errorf("рейтинг для этого пользователя уже существует")
	}

	// Создаем новый рейтинг
	newRating := &models.Rating{
		VideoID: videoID,
		UserID:  userID,
		Rating:  rating,
	}

	err = s.RatingRepository.AddRating(newRating)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении рейтинга: %v", err)
	}

	return nil
}

// GetAverageRatingForVideo возвращает средний рейтинг для указанного видео
func (s *RatingService) GetAverageRatingForVideo(videoID int) (float64, error) {
	avgRating, err := s.RatingRepository.GetAverageRatingForVideo(videoID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при получении среднего рейтинга: %v", err)
	}
	return avgRating, nil
}
