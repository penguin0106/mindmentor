package models

// Emotion представляет собой структуру данных для эмоции (записи)
type Emotion struct {
	ID     int    `json:"id"`
	Topic  string `json:"topic"`
	Body   string `json:"body"`
	UserID int    `json:"userID"`
}
