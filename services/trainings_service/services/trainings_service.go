package services

import (
	"errors"
	"mindmentor/services/trainings_service/models"
	"mindmentor/services/trainings_service/repositories"
)

type TrainingService struct {
	TrainRepo *repositories.TrainingRepository
}

func NewTrainingService(trainRepo *repositories.TrainingRepository) *TrainingService {
	return &TrainingService{TrainRepo: trainRepo}
}

// GetAllTrainings возвращает все тренировки
func (s *TrainingService) GetAllTrainings() ([]*models.Training, error) {
	trainings, err := s.TrainRepo.GetAllTrainings()
	if err != nil {
		return nil, err
	}
	return trainings, nil
}

func (s *TrainingService) GetTrainingByName(trainingName string) (*models.Training, error) {
	training, err := s.TrainRepo.GetTrainingByName(trainingName)
	if err != nil {
		return nil, err
	}
	if training == nil {
		return nil, errors.New("тренировка не найдена")
	}
	return training, nil
}

// AddTraining добавляет новую тренировку
func (s *TrainingService) AddTraining(training *models.Training) error {
	err := s.TrainRepo.AddTraining(training)
	if err != nil {
		return errors.New("ошибка при добавлении тренировки")
	}
	return nil
}
