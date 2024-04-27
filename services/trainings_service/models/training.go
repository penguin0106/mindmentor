package models

// Training представляет модель данных для тренировки
type Training struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Rating      float64  `json:"rating"`
	Favorite    bool     `json:"favorite"`
	Comments    []string `json:"comments"`
}
