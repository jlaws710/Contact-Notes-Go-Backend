package model

type Note struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ContactID uint   `json:"contact_id"`
	Body      string `json:"body"`
}
