package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" binding:"required,max=50"`
	Content string `gorm:"not null" binding:"required,max=2000"`
	UserID  uint
	User    User `binding:"omitempty"`
}
