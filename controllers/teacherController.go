package controllers

import (
	"net/http"
	"github.com/HAGG-glitch/student_managment_system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateTeacher handles the creation of a new teacher
func CreateTeacher (c *gin.Context, db *gorm.DB) {
	var teacher models.Teacher
	// Validate Input
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// Attempt to create teacher in the database
	if err := db.Create(&teacher).Error; err != nil{
		// If there is any error (like duplicate employee_id), return it
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Successfully created
	c.IndentedJSON(http.StatusCreated, teacher)
}

// GetTeachers retrieves all teachers from the database
func GetTeachers (c *gin.Context, db *gorm.DB) {
	var teachers []models.Teacher
	if err := db.Preload("Class").Find(&teachers).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, teachers)

	if len(teachers) == 0 {
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No teachers found"})
    return
}
}

// Get Teacher by ID
func GetTeacherByID(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	if err := db.Find(&models.Teacher{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, models.Teacher{})
}

// Update Teacher by ID
func UpdateTeacher(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	var teacher models.Teacher
	// Validate input
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Attempt to update teacher in the database
	if err := db.Model(&models.Teacher{}).Where("id = ?", id).Updates(teacher).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Teacher updated successfully"})
}


// Delete Teacher by ID
func DeleteTeacher(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	// Attempt to delete teacher from the database
	if err := db.Delete(&models.Teacher{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}