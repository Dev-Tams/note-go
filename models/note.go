package models

type Note struct {
	ID    uint   `gorm:"id"`
	Title  string `json:"title"`
	Content string `json:"content"`
	User_id uint `json:"user_id"`
	// User User `json:"user"`
}