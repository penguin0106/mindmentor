package services

import (
	"meditation_service/models"
	"meditation_service/repositories"
	"time"
)

type CommentService struct {
	CommentRepository *repositories.CommentRepository
}

func NewCommentService(commentRepo *repositories.CommentRepository) *CommentService {
	return &CommentService{CommentRepository: commentRepo}
}

// AddVideoComment добавляет новый комментарий курса медитации
func (s *CommentService) AddVideoComment(userID, itemID int, text string) error {
	comment := &models.Comment{
		UserID:    userID,
		ItemID:    itemID,
		Text:      text,
		Timestamp: time.Now().Unix(),
	}

	err := s.CommentRepository.AddComment(comment)
	if err != nil {
		return err
	}
	return nil
}

// GetVideoComments возвращает все комментарии для указанного курса медитации
func (s *CommentService) GetVideoComments(videoID int) ([]*models.Comment, error) {
	comments, err := s.CommentRepository.GetCommentsByVideoID(videoID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
