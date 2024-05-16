package services

import (
	"trainings_service/models"
	"trainings_service/repositories"
)

type TrainingService struct {
	TrainRepo *repositories.TrainingRepository
}

func NewTrainingService(trainRepo *repositories.TrainingRepository) *TrainingService {
	return &TrainingService{TrainRepo: trainRepo}
}

// AddBook добавляет новую книгу
func (s *TrainingService) AddBook(title, description string, content []byte) error {
	book := &models.Book{
		Title:       title,
		Description: description,
		Content:     content,
	}

	err := s.TrainRepo.AddBook(book)
	if err != nil {
		return err
	}
	return nil
}

// GetBookByID возвращает книгу по ее идентификатору
func (s *TrainingService) GetBookByID(bookID int) (*models.Book, error) {
	book, err := s.TrainRepo.GetBookByID(bookID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

// GetAllBooks возвращает все книги
func (s *TrainingService) GetAllBooks() ([]*models.Book, error) {
	books, err := s.TrainRepo.GetAllBook()
	if err != nil {
		return nil, err
	}
	return books, nil
}

// GetBookByName возвращает книгу по ее названию
func (s *TrainingService) GetBookByName(title string) (*models.Book, error) {
	book, err := s.TrainRepo.GetBookByName(title)
	if err != nil {
		return nil, err
	}
	return book, nil
}
