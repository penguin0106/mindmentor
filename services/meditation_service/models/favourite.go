package models

// Favorite представляет собой элемент, добавленный в избранное
type Favorite struct {
	UserID int `json:"user_id"`
	ItemID int `json:"item_id"`
	// Другие поля, например, тип элемента (курс или музыка)
}
