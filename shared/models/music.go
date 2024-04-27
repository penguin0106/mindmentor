package models

// Music представляет аудиофайл для медитации
type Music struct {
	ID       int    // Уникальный идентификатор аудиофайла
	Name     string // Название аудиофайла
	Duration int    // Продолжительность аудиофайла в секундах
	URL      string // Ссылка на аудиофайл
}
