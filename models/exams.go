package models
import "gorm.io/gorm"

type Exam struct {
	gorm.Model
	StudentID uint    `json:"student_id"`
	Student   Student `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubjectID uint    `json:"subject_id"`
	Subject   Subject `gorm:"foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Score     int     `json:"score" binding:"required"`
	Grade     string  `json:"grade" binding:"required"`
}
