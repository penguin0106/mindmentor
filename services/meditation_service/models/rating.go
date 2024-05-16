package models

// Rating представляет собой рейтинг видео
type Rating struct {
	ID        int     `json:"id"`
	VideoID   int     `json:"video_id"`
	UserID    int     `json:"user_id"`
	Rating    float64 `json:"rating"`
	CreatedAt string  `json:"created_at"`
}
