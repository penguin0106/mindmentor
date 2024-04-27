package models

type Discussion struct {
	ID      int    `json:"id"`
	Topic   string `json:"topic"`
	OwnerID int    `json:"ownerId"`
}
