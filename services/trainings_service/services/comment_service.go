package services

import (
	"errors"
	"mindmentor/services/trainings_service/repositories"
	"mindmentor/shared/models"
)

type CommentService struct {
	ComRepo *repositories.CommentRepository
}

func NewCommentService(comRepo *repositories.CommentRepository) *CommentService {
	return &CommentService{ComRepo: comRepo}
}

// AddComment добавляет новый комментарий к тренировке
func (s *CommentService) AddComment(userID, trainingID int, text string) error {
	err := s.ComRepo.AddComment(userID, trainingID, text)
	if err != nil {
		return errors.New("ошибка при добавлении комментария")
	}
	return nil
}

// GetCommentsByTrainingID возвращает все комментарии для указанной тренировки
func (s *CommentService) GetCommentsByTrainingID(trainingID int) ([]*models.Comment, error) {
	comments, err := s.ComRepo.GetCommentsByTrainingID(trainingID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// AddRating добавляет оценку для указанной тренировки
func (s *CommentService) AddRating(rating *models.Rating) error {
	err := s.ComRepo.AddRating(rating)
	if err != nil {
		return errors.New("ошибка при добавлении оценки")
	}
	return nil
}

// GetRating возвращает рейтинг тренировки по ее идентификатору
func (s *CommentService) GetRating(trainingID int) (float64, error) {
	rating, err := s.ComRepo.GetRating(trainingID)
	if err != nil {
		return 0, err
	}
	return rating, nil
}
