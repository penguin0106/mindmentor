package models

// Favorite представляет собой элемент, добавленный в избранное
type Favorite struct {
	UserID int `json:"userId"`
	ItemID int `json:"itemId"`
	// Другие поля, например, тип элемента (курс или музыка)
}
