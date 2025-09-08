package controllers

import (
	"net/http"
	"github.com/HAGG-glitch/student_managment_system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateClass handles the creation of a new class
func CreateClass (c *gin.Context, db *gorm.DB) {
	var class models.Class

	// validate the input
	if err := c.ShouldBindJSON(&class); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}


	// attempt to create class in the database
	if err := db.Create(&class).Error; err != nil {
		// If there is any error (like duplicate name), return it
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Successfully created
	db.Create(&class)
	c.IndentedJSON(http.StatusCreated, class)
}


// GetClasses retrieves all classes from the database
func GetClasses (c *gin.Context, db *gorm.DB) {
	var classes []models.Class

	

	if err := db.Preload("Students").Preload("Teachers").Find(&classes).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	if len(classes) == 0 {
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No teachers found"})
    return
}

c.IndentedJSON(http.StatusOK, classes)

}

// Get Class by ID
func GetClassByID(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	if err := db.Find(&models.Class{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, models.Class{})
}




// Update Class by ID
func UpdateClass(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	var class models.Class
	// Validate input
	if err := c.ShouldBindJSON(&class); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Attempt to update class in the database
	if err := db.Model(&models.Class{}).Where("id = ?", id).Updates(class).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Class updated successfully"})
}


// Delete class by ID
func DeleteClass(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	// Attempt to delete class from the database
	if err := db.Delete(&models.Class{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}