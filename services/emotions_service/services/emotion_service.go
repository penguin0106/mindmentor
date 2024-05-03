package services

import (
	"mindmentor/services/emotions_service/models"
	"mindmentor/services/emotions_service/repositories"
)

type EmotionService struct {
	EmotionRepository *repositories.EmotionRepository
}

func NewEmotionService(emoRepo *repositories.EmotionRepository) *EmotionService {
	return &EmotionService{EmotionRepository: emoRepo}
}

// CreateEmotion создает новую запись эмоции
func (s *EmotionService) CreateEmotion(emotion *models.Emotion) error {
	return s.EmotionRepository.CreateEmotion(emotion)
}

// UpdateEmotion обновляет существующую запись эмоции
func (s *EmotionService) UpdateEmotion(emotionID int, updatedEmotion *models.Emotion) error {
	return s.EmotionRepository.UpdateEmotion(emotionID, updatedEmotion)
}

// DeleteEmotion удаляет запись эмоции
func (s *EmotionService) DeleteEmotion(emotionID int) error {
	return s.EmotionRepository.DeleteEmotion(emotionID)
}

// GetEmotionsByUserID возвращает эмоции пользователя по его идентификатору
func (s *EmotionService) GetEmotionsByUserID(userID int) ([]*models.Emotion, error) {
	return s.EmotionRepository.GetEmotionsByUserID(userID)
}
