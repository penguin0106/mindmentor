package repositories

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"mindmentor/shared/models"
)

// TokenRepository представляет собой репозиторий для работы с токенами доступа
type TokenRepository struct {
	DB *sql.DB
}

// GetTokenByUserID получает токен из базы данных по идентификатору пользователя
func (r *TokenRepository) GetTokenByUserID(userID int) (*models.Token, error) {
	query := "SELECT token FROM tokens WHERE user_id = $1"
	row := r.DB.QueryRow(query, userID)

	var token string
	err := row.Scan(&token)
	if err != nil {
		return nil, err
	}

	return &models.Token{Token: token}, nil
}

// CheckTokenValidity проверяет валидность токена доступа
func (r *TokenRepository) CheckTokenValidity(token string) (bool, error) {
	query := "SELECT COUNT(*) FROM tokens WHERE token = $1"
	var count int
	err := r.DB.QueryRow(query, token).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// GenerateToken генерирует и сохраняет новый токен доступа для указанного пользователя
func (r *TokenRepository) GenerateToken(userID int) (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	token := hex.EncodeToString(tokenBytes)

	query := "INSERT INTO tokens (token, user_id) VALUES ($1, $2)"
	_, err = r.DB.Exec(query, token, userID)
	if err != nil {
		return "", err
	}

	return token, nil
}
