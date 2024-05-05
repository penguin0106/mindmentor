package services

import (
	"errors"
	"trainings_service/repositories"
)

type FavoriteService struct {
	FavRepo *repositories.FavoriteRepository
}

func NewFavoriteService(favRepo *repositories.FavoriteRepository) *FavoriteService {
	return &FavoriteService{FavRepo: favRepo}
}

// AddToFavorites добавляет тренировку в избранное для указанного пользователя
func (s *FavoriteService) AddToFavorites(userID, trainingID int) error {
	err := s.FavRepo.AddToFavorites(userID, trainingID)
	if err != nil {
		return errors.New("ошибка при добавлении тренировки в избранное")
	}
	return nil
}

// RemoveFromFavorites удаляет тренировку из избранного для указанного пользователя
func (s *FavoriteService) RemoveFromFavorites(userID, trainingID int) error {
	err := s.FavRepo.RemoveFromFavorites(userID, trainingID)
	if err != nil {
		return errors.New("ошибка при удалении тренировки из избранного")
	}
	return nil
}
