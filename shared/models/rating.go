package models

// Rating представляет оценку курса медитации
type Rating struct {
	ID     int     // Идентификатор оценки
	ItemID int     // Идентификатор курса
	UserID int     // Идентификатор пользователя, оставившего оценку
	Value  float64 // Значение оценки
}
