package models

import "time"

// Music представляет аудиофайл для медитации
type Music struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Duration  int       `json:"duration"`
	MusicFile []byte    `json:"music_file"`
	CreatedAt time.Time `json:"created_at"`
}
