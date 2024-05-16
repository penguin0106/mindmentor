package services

import (
	"fmt"
	"meditation_service/models"
	"meditation_service/repositories"
)

type FavoriteService struct {
	FavoriteRepo *repositories.FavoriteRepository
}

func NewFavoriteService(favRepo *repositories.FavoriteRepository) *FavoriteService {
	return &FavoriteService{FavoriteRepo: favRepo}
}

func (s *FavoriteService) AddToFavourite(fav *models.Favorite) error {
	err := s.FavoriteRepo.AddToFavorite(fav)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении видео в избранное: %v", err)
	}
	return nil
}

func (s *FavoriteService) RemoveFromFavorite(fav *models.Favorite) error {
	err := s.FavoriteRepo.RemoveFromFavorite(fav)
	if err != nil {
		return fmt.Errorf("ошибка при удалении видео из избранного: %v", err)
	}
	return nil
}
