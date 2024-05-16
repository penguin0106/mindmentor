package models

// User представляет собой модель пользователя
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Другие поля профиля пользователя, если необходимо
}

// Favorite представляет собой модель избранного элемента
type Favorite struct {
	UserID int `json:"user_id"`
	ItemID int `json:"item_id"`
	// Другие поля, если необходимо, например, тип элемента (видео, тренировка)
}
