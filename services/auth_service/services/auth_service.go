package services

import (
	"auth_service/models"
	"auth_service/repositories"
	"errors"
	"regexp"
)

// AuthService представляет сервис аутентификации и регистрации пользователей
type AuthService struct {
	UserRepository *repositories.UserRepository
}

// NewAuthService создает новый экземпляр сервиса аутентификации
func NewAuthService(authRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: authRepo,
	}
}

// RegisterUser регистрирует нового пользователя
func (svc *AuthService) RegisterUser(username, email, password string) error {
	if err := svc.validatePassword(password); err != nil {
		return err
	}
	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
	}
	err := svc.UserRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}

// AuthenticateUser аутентифицирует пользователя по его учетным данным
func (svc *AuthService) AuthenticateUser(identifier, password string) (*models.User, error) {
	user, err := svc.UserRepository.Authenticate(identifier, password)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("пользователь не найден")
	}
	return user, nil
}

// validatePassword проверяет валидность пароля
func (svc *AuthService) validatePassword(password string) error {
	// Проверяем длину пароля
	if len(password) == 0 {
		return errors.New("password cannot be empty")
	}
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Проверяем наличие хотя бы одной цифры
	hasDigit, err := regexp.MatchString(`\d`, password)
	if err != nil {
		return errors.New("error while validating password")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}

	// Проверяем наличие хотя бы одной буквы в верхнем регистре
	hasUpperCase, err := regexp.MatchString(`[A-Z]`, password)
	if err != nil {
		return errors.New("error while validating password")
	}
	if !hasUpperCase {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Проверяем наличие хотя бы одной буквы в нижнем регистре
	hasLowerCase, err := regexp.MatchString(`[a-z]`, password)
	if err != nil {
		return errors.New("error while validating password")
	}
	if !hasLowerCase {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Проверяем наличие хотя бы одного специального символа
	hasSpecialChar, err := regexp.MatchString(`[^A-Za-z0-9]`, password)
	if err != nil {
		return errors.New("error while validating password")
	}
	if !hasSpecialChar {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
