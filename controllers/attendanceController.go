package controllers

import (
	"net/http"
	"time"

	"github.com/HAGG-glitch/student_managment_system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MarkAttendance marks attendance for a student in a subject
func MarkAttendance(c *gin.Context, db *gorm.DB) {
	var input struct {
		StudentID uint      `json:"student_id" binding:"required"`
		ClassID   uint      `json:"class_id" binding:"required"`
		Status    string    `json:"status" binding:"required,oneof=present absent late"`
		Date      time.Time `json:"date" binding:"-"`
	}


	// validate the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// attempt to create class in the database
	attendance := models.Attendance{
		StudentID: input.StudentID,
		ClassID:   input.ClassID,
		Status:    input.Status,
		Date: input.Date,
	}
	if err := db.Create(&attendance).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, attendance)
}

// GetAttendance retrieves attendance records for a specific student
func GetAttendanceByStudent(c *gin.Context, db *gorm.DB) {
	studentID := c.Param("student_id")
	var attendance []models.Attendance
	// searching for the student attendance
	if err := db.Where("student_id = ?", studentID).Preload("Student").Preload("Class").Find(&attendance).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(attendance) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No attendance records found for the student"})
		return
	}
	c.IndentedJSON(http.StatusOK, attendance)
}

// GetAttendanceByClass retrieves attendance records for a specific class
func GetAttendanceByClass(c *gin.Context, db *gorm.DB) {
	classID := c.Param("class_id")
	var attendance []models.Attendance
	// searching for the student attendance
	if err := db.Where("class_id = ?", classID).Preload("Student").Preload("Class").Find(&attendance).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(attendance) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No attendance records found for the student"})
		return
	}
	c.IndentedJSON(http.StatusOK, attendance)
}

// UpdateAttendance updates an attendance record by ID
func UpdateAttendance(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var attendance models.Attendance
	// Validate input
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// attempt to update attendance in the database
	if err := db.Model(&models.Attendance{}).Where("id = ?", id).Updates(attendance).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Attendance updated successfully"})
}


