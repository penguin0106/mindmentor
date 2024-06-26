package models

// Comment представляет собой комментарий пользователя
type Comment struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	ItemID    int    `json:"itemId"`
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"` // Временная метка комментария
	// Другие поля, если нужны
}
