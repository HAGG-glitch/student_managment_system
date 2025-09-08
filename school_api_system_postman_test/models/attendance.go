package models

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	StudentID uint      `json:"student_id"`
	Student   Student   `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ClassID   uint      `json:"class_id"`
	Class     Class     `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status" binding:"required" enums:"Present,Absent,Late"` // Present, Absent, Late
}