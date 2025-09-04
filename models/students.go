package models

import "gorm.io/gorm"


type Student struct {
	gorm.Model
	AdmissionID string  `json:"admission_id" gorm:"unique;not null" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Age         int     `json:"age" binding:"required"`
	Email       string  `json:"email" gorm:"unique;not null" binding:"required"`
	ClassID     *uint   `json:"class_id"` // Nullable foreign key
	Class       *Class  `json:"class" gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Enrollments []Enrollment
	Exams       []Exam
	Attendance  []Attendance
}
