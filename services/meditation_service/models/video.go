package models

import "time"

// Video представляет собой курс медитации
type Video struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	VideoContent []byte    `json:"video_content"`
	CreatedAt    time.Time `json:"created_at"`
	// Другие поля, например, автор, длительность, уровень сложности и т.д.
}
