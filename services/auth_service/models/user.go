package models

// User структура представляет модель данных пользователя
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" unique:"true"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
