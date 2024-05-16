package services

import (
	"errors"
	"profile_service/models"
	"profile_service/repositories"
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

// EditProfileUsername редактирует имя пользователя в профиле
func (s *UserService) EditProfileUsername(userID int, newUsername string) error {
	err := s.userRepository.EditProfileUsername(userID, newUsername)
	if err != nil {
		return errors.New("ошибка при редактировании имени пользователя")
	}
	return nil
}

// EditProfileEmail редактирует email в профиле
func (s *UserService) EditProfileEmail(userID int, newEmail string) error {
	err := s.userRepository.EditProfileEmail(userID, newEmail)
	if err != nil {
		return errors.New("ошибка при редактировании email пользователя")
	}
	return nil
}

// EditProfilePassword редактирует пароль в профиле
func (s *UserService) EditProfilePassword(userID int, newPassword string) error {
	err := s.userRepository.EditProfilePassword(userID, newPassword)
	if err != nil {
		return errors.New("ошибка при редактировании пароля пользователя")
	}
	return nil
}

// GetFavoriteVideos возвращает избранные видео пользователя
func (s *UserService) GetFavoriteVideos(userID int) ([]int, error) {
	favoriteVideos, err := s.userRepository.GetFavoriteVideos(userID)
	if err != nil {
		return nil, err
	}
	return favoriteVideos, nil
}

// GetFavoriteTrainings возвращает избранные тренировки пользователя
func (s *UserService) GetFavoriteTrainings(userID int) ([]int, error) {
	favoriteTrainings, err := s.userRepository.GetFavoriteTrainings(userID)
	if err != nil {
		return nil, err
	}
	return favoriteTrainings, nil
}
