package models

type Activity struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	Steps      int     `json:"steps"`
	SleepHours float64 `json:"sleep_hours"`
}
