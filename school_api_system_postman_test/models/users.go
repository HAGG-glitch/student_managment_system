package models

import "gorm.io/gorm"



type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" gorm:"unique;not null" binding:"required"`
	Role     string `json:"role" binding:"required" enums:"admin,teacher,student"` // e.g., "admin", "teacher", "student"
}
