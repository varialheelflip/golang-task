package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" binding:"required,max=20"`
	Password string `gorm:"not null" binding:"required,max=20"`
	Email    string `gorm:"unique;not null" binding:"required,email,max=30"`
}
