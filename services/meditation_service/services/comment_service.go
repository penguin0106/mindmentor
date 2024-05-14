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

// AddCourseComment добавляет новый комментарий курса медитации
func (s *CommentService) AddCourseComment(userID, itemID int, text string) error {
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

// GetCourseComments возвращает все комментарии для указанного курса медитации
func (s *CommentService) GetCourseComments(courseID int) ([]*models.Comment, error) {
	comments, err := s.CommentRepository.GetCommentsByCourseID(courseID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
