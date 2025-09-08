package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	EmployeeID string  `json:"employee_id" gorm:"unique;not null" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Age        int     `json:"age" binding:"required"`
	Email      string  `json:"email" gorm:"unique;not null" binding:"required"`
	ClassID    *uint   `json:"class_id"`
	Class      *Class  `json:"class" gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Subjects   []Subject
}