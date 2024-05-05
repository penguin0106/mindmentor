package services

import (
	"meditation_service/models"
	"meditation_service/repositories"
)

type CommentService struct {
	CommentRepository *repositories.CommentRepository
}

func NewCommentService(commentRepo *repositories.CommentRepository) *CommentService {
	return &CommentService{CommentRepository: commentRepo}
}

// AddCourseComment добавляет новый комментарий курса медитации
func (s *CommentService) AddCourseComment(comment *models.Comment) error {
	return s.CommentRepository.AddComment(comment)
}

// GetCourseComments возвращает все комментарии для указанного курса медитации
func (s *CommentService) GetCourseComments(courseID int) ([]*models.Comment, error) {
	return s.CommentRepository.GetCommentsByCourseID(courseID)
}
