package models

type Meditation struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VideoURL    string `json:"video_url"`
	AudioURL    string `json:"audio_url"`
}
