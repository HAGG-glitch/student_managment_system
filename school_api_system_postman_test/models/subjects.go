package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	SubjectCode string   `json:"subject_code" gorm:"unique;not null" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	TeacherID   uint     `json:"teacher_id"`
	Teacher     Teacher  `json:"teacher" gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Credits     int      `json:"credits" binding:"required"`
	Hours       int      `json:"hours" binding:"required"`
	Enrollments []Enrollment
	Exams       []Exam
}