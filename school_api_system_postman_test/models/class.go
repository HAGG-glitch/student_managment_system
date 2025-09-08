package models

import "gorm.io/gorm"


type Class struct {
	gorm.Model
	Name     string    `json:"name" binding:"required"`
	Students  []Student
	Subjects  []Subject `gorm:"many2many:class_subjects"`
}