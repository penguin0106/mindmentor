package models

type Chat struct {
	ID    int    `json:"id"`
	Topic string `json:"topic"`
}

type ChatParticipant struct {
	ID     int `json:"id"`
	ChatID int `json:"chat_id"`
	UserID int `json:"user_id"`
}
