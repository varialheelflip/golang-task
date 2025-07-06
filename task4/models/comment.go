package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" binding:"required,max=255"`
	UserID  uint
	User    User `binding:"omitempty"`
	PostID  uint `binding:"required"`
	Post    Post `binding:"omitempty"`
}
