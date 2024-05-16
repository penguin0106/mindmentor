package services

import (
	"auth_service/repositories"
	"fmt"
)

// JWTService представляет сервис для работы с JWT токенами
type JWTService struct {
	JWTRepository *repositories.JWTRepository
}

// NewJWTService создает новый экземпляр сервиса для работы с JWT токенами
func NewJWTService(jwtRepo *repositories.JWTRepository) *JWTService {
	return &JWTService{JWTRepository: jwtRepo}
}

// GenerateToken генерирует новый JWT токен для пользователя с указанным идентификатором
func (s *JWTService) GenerateToken(userID int, username, password string) (string, error) {
	token, err := s.JWTRepository.GenerateToken(userID, username, password)
	if err != nil {
		return "", fmt.Errorf("ошибка при генерации токена: %v", err)
	}
	return token, nil
}

// VerifyToken проверяет действительность JWT токена и возвращает данные пользователя
func (s *JWTService) VerifyToken(token string) (int, string, string, error) {
	userID, username, password, err := s.JWTRepository.VerifyToken(token)
	if err != nil {
		return 0, "", "", fmt.Errorf("ошибка при верификации токена: %v", err)
	}
	return userID, username, password, nil
}
