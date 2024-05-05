package services

import "meditation_service/repositories"

type FavoriteService struct {
	FavoriteRepo *repositories.FavoriteRepository
}

func NewFavoriteService(favRepo *repositories.FavoriteRepository) *FavoriteService {
	return &FavoriteService{FavoriteRepo: favRepo}
}

// AddToFavorite добавляет элемент в избранное для указанного пользователя
func (s *FavoriteService) AddToFavorite(userID, itemID int) error {
	return s.FavoriteRepo.AddToFavorite(userID, itemID)
}

// RemoveFromFavorite удаляет элемент из избранного для указанного пользователя
func (s *FavoriteService) RemoveFromFavorite(userID, itemID int) error {
	return s.FavoriteRepo.RemoveFromFavorite(userID, itemID)
}
