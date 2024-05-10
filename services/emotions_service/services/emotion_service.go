package services

import (
	"emotions_service/models"
	"emotions_service/repositories"
)

type EmotionService struct {
	EmotionRepository *repositories.EmotionRepository
}

func NewEmotionService(emoRepo *repositories.EmotionRepository) *EmotionService {
	return &EmotionService{EmotionRepository: emoRepo}
}

// CreateEmotion создает новую запись эмоции
func (s *EmotionService) CreateEmotion(topic, body string) error {
	emotion := &models.Emotion{
		Topic: topic,
		Body:  body,
	}
	err := s.EmotionRepository.CreateEmotion(emotion)
	if err != nil {
		return err
	}
	return nil
}

// UpdateEmotion обновляет существующую запись эмоции
func (s *EmotionService) UpdateEmotion(emotionID int, updatedEmotion *models.Emotion) error {
	return s.EmotionRepository.UpdateEmotion(emotionID, updatedEmotion)
}

// DeleteEmotion удаляет запись эмоции
func (s *EmotionService) DeleteEmotion(emotionID int) error {
	err := s.EmotionRepository.DeleteEmotion(emotionID)
	if err != nil {
		return err
	}
	return nil
}

// GetEmotionsByUserID возвращает эмоции пользователя по его идентификатору
func (s *EmotionService) GetEmotionsByUserID() ([]*models.Emotion, error) {
	emotion, err := s.EmotionRepository.GetEmotionsByUserID()
	if err != nil {
		return nil, err
	}
	return emotion, nil
}
