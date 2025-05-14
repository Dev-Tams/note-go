package models

type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email" binding:"required" gorm:"unique"`
}
