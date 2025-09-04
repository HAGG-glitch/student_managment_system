package models

import(
	"gorm.io/gorm"
)

type Enrollment struct {
	gorm.Model
	StudentID uint    `json:"student_id"`
	Student   Student `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubjectID uint    `json:"subject_id"`
	Subject   Subject `gorm:"foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Grade     string  `json:"grade" binding:"required"`
}
