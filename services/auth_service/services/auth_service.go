package handlers

import (
	"errors"
	"mindmentor/services/auth_service/repositories"
	"mindmentor/shared/models"
	"regexp"
)

// AuthService представляет сервис авторизации
type AuthService struct {
	UserRepository *repositories.UserRepository
}

// NewAuthService создает новый экземпляр сервиса авторизации
func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepo}
}

// RegisterUser регистрирует нового пользователя
func (service *AuthService) RegisterUser(user *models.User) error {
	// Проверка валидности пароля
	if err := service.validatePassword(user.Password); err != nil {
		return err
	}

	// Проверка, что пользователь с таким email не существует
	existingUser, err := service.UserRepository.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	// Сохранение пользователя в базе данных
	if err := service.UserRepository.Save(user); err != nil {
		return err
	}
	return nil
}

// Authenticate аутентифицирует пользователя по email и паролю
func (service *AuthService) Authenticate(email, password string) (bool, error) {
	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil // Пользователь с таким email не найден
	}
	// Проверка пароля
	if user.Password != password {
		return false, nil // Неправильный пароль
	}
	return true, nil // Аутентификация успешна
}

// validatePassword проверяет валидность пароля
func (service *AuthService) validatePassword(password string) error {
	// Проверяем длину пароля
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
