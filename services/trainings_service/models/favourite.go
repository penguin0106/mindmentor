package models

// Favorite представляет модель данных для избранных тренировок пользователя
type Favorite struct {
	UserID     int `json:"userId"`
	TrainingID int `json:"trainingId"`
}
