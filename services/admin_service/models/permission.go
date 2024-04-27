package models

// Permission представляет модель разрешения
type Permission struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
