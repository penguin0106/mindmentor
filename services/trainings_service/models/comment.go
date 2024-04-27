package models

// Comment представляет модель данных для комментария
type Comment struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
}
