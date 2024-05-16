package models

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	Content     []byte
	CreatedAt   time.Time
}
