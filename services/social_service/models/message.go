package models

import "time"

// Message представляет сообщение в обсуждении.
type Message struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	DiscussionID int       `json:"discussion_id"`
	Text         string    `json:"text"`
	CreationTime time.Time `json:"creation_time"`
	LastEditTime time.Time `json:"last_edit_time"`
}
