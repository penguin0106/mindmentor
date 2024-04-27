package models

// User структура представляет модель данных пользователя
type User struct {
	ID       int
	Login    string
	Password string
	Email    string
	// Другие поля пользователя
}
