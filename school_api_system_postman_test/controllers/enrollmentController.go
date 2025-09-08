package controllers

import (
	"net/http"
	"github.com/HAGG-glitch/student_managment_system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEnrollment(c *gin.Context, db *gorm.DB) {
	var enroll models.Enrollment

	// Validate input
	if err := c.ShouldBindJSON(&enroll); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// attempt to create exam in the databse
	if err := db.Create(&enroll).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	


	// successfully created
	c.IndentedJSON(http.StatusCreated, enroll)
}

func GetEnrollments(c *gin.Context, db *gorm.DB) {
	var enrolls []models.Enrollment

	if err := db.Preload("Subject").Preload("Student").Find(&enrolls).Error; err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(enrolls) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No enrollments were made"})
		return
	}

	c.IndentedJSON(http.StatusOK, enrolls)
}


// Find by ID
func GetEnrollmentByID(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	if err := db.Find(&models.Enrollment{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Enrollment not found"} )

}



// UpdateStudent updates a student's information by ID
func UpdateEnrollment(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	var enroll models.Enrollment
	// Validate input
	if err := c.ShouldBindJSON(&enroll); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// Attempt to update student in the database
	if err := db.Model(&models.Exam{}).Where("id = ?", id).Updates(enroll).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Successfully updated
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Enrollment updated successfully"})
}



// DeleteStudent deletes a student by ID
func DeleteEnrollment(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	if err := db.Delete(&models.Enrollment{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Enrollment is deleted successfully"})

}