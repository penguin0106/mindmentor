package services

import (
	"trainings_service/models"
	"trainings_service/repositories"
)

// CommentService представляет сервис для работы с комментариями к книгам
type CommentService struct {
	CommentRepo *repositories.CommentRepository
}

// NewCommentService создает новый экземпляр сервиса для работы с комментариями
func NewCommentService(commentRepo *repositories.CommentRepository) *CommentService {
	return &CommentService{CommentRepo: commentRepo}
}

// AddBookComment добавляет новый комментарий к книге
func (s *CommentService) AddBookComment(userID, bookID int, text string) error {
	err := s.CommentRepo.AddComment(userID, bookID, text)
	if err != nil {
		return err
	}
	return nil
}

// GetBookComments возвращает все комментарии для указанной книги
func (s *CommentService) GetBookComments(bookID int) ([]*models.Comment, error) {
	comments, err := s.CommentRepo.GetCommentsByBookID(bookID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// AddBookRating добавляет оценку для указанной книги
func (s *CommentService) AddBookRating(rating *models.Rating) error {
	err := s.CommentRepo.AddRating(rating)
	if err != nil {
		return err
	}
	return nil
}

// GetBookRating возвращает рейтинг книги по ее идентификатору
func (s *CommentService) GetBookRating(bookID int) (float64, error) {
	rating, err := s.CommentRepo.GetRating(bookID)
	if err != nil {
		return 0, err
	}
	return rating, nil
}
