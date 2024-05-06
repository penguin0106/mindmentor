package models

// TokenClaims представляет структуру токена JWT
type TokenClaims struct {
	Exp    int64  `json:"exp"`     // Срок действия токена (timestamp)
	UserID string `json:"user_id"` // Идентификатор пользователя
}
