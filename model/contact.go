package model

type Contact struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Notes    []Note `gorm:"constraint:OnDelete:CASCADE"`
}
