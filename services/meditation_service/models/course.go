package models

// Course представляет собой курс медитации
type Course struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Другие поля, например, автор, длительность, уровень сложности и т.д.
}
