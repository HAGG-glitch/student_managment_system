package controllers

import (
	"net/http"

	"github.com/HAGG-glitch/student_managment_system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateSubject handles the creation of a new subject

func CreateSubject (c *gin.Context, db *gorm.DB) {
	var subject models.Subject
	// Validate input
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// attempt to create subject in the databse
	if err := db.Create(&subject).Error; err != nil{
		// If there is any error (like duplicate subject code), return it
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Successfully created
	c.IndentedJSON(http.StatusCreated, subject)
}

// GetSubjects retrieves all subjects from the database
func GetSubjects (c *gin.Context, db *gorm.DB) {
	var subjects []models.Subject

	
	
	if err := db.Find(&subjects).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(subjects) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No subjects found"})
		return
	}
	c.IndentedJSON(http.StatusOK, subjects)

}


// Get Subject by ID
func GetSubjectByID(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	if err := db.Find(&models.Subject{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, models.Subject{})
}




// Update Subject by ID
func UpdateSubject(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	var subject models.Subject
	// Validate input
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Attempt to update subject in the database
	if err := db.Model(&models.Subject{}).Where("id = ?", id).Updates(subject).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Subject updated successfully"})
}


// Delete subject by ID
func DeleteSubject(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	// Attempt to delete subject from the database
	if err := db.Delete(&models.Subject{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
}

