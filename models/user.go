package models

type User struct {
	ID    uint   `gorm:"id"`
	Name  string `json:"name"`
	Email string `json:"email" binding:"required"`
}
