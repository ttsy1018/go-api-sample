package models

type Todo struct {
	BaseModel
	Title string `gorm:"size:255" json:"title,omitempty"`
	Comment string `gorm:"type:text" json:"text,omitempty"`
	UserID int `gorm:"not null" json:"user_id"`
}