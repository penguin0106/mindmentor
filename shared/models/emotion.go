package models

import "time"

// Emotion представляет собой структуру данных для эмоции (записи)
type Emotion struct {
	ID        int       `json:"id"`
	Topic     string    `json:"topic"`
	Body      string    `json:"body"`
	UserID    int       `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
