package services

import (
	"mindmentor/services/profile_service/repositories"
	"mindmentor/shared/models"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

// GetUserByID возвращает пользователя по его идентификатору
func (s *UserService) GetUserByID(userID int) (*models.User, error) {
	// Используем метод GetUserByID из userRepository для получения пользователя
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser обновляет информацию о пользователе
func (s *UserService) UpdateUser(userID int, updatedUser *models.User) error {
	// Используем метод UpdateUser из userRepository для обновления пользователя
	err := s.userRepository.UpdateUser(userID, updatedUser)
	if err != nil {
		return err
	}
	return nil
}

// GetFavoriteCourse возвращает избранные курсы для указанного пользователя
func (s *UserService) GetFavoriteCourse(userID int) ([]models.Favorite, error) {
	return s.userRepository.GetFavoriteCourse(userID)
}

// GetFavoriteTraining возвращает избранные тренировки для указанного пользователя
func (s *UserService) GetFavoriteTraining(userID int) ([]models.Favorite, error) {
	return s.userRepository.GetFavoriteTraining(userID)
}
