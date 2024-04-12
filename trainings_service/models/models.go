package models

type Training struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}

type Review struct {
	ID          int    `json:"id"`
	TrainingID  int    `json:"training_id"`
	UserID      int    `json:"user_id"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}
