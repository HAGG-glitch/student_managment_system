package controllers

import (
	"net/http"
	"github.com/HAGG-glitch/student_managment_system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateStudent handles the creation of a new student
func CreateStudent (c *gin.Context, db *gorm.DB) {
	var student models.Student
	// Validate input
	if err := c.ShouldBindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	} 

	// attempt to create student in the databse
	if err := db.Create(&student).Error; err != nil{
		// If there is any error (like duplicate admission_id), return it
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Successfully created
	c.IndentedJSON(http.StatusCreated, student)

}


// GetStudents retrieves all students from the database
func GetStudents (c *gin.Context, db *gorm.DB) {
	var students []models.Student
	
	

	if err := db.Debug().Preload("Class").Find(&students).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(students) == 0 {
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No students found"})
    return
	}

	c.IndentedJSON(http.StatusOK, students)

	
	
}

// Find by ID
func GetStudentByID(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	if err := db.Find(&models.Student{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student found"} )

}



// UpdateStudent updates a student's information by ID
func UpdateStudent(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	var student models.Student
	// Validate input
	if err := c.ShouldBindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// Attempt to update student in the database
	if err := db.Model(&models.Student{}).Where("id = ?", id).Updates(student).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Successfully updated
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}



// DeleteStudent deletes a student by ID
func DeleteStudent(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	if err := db.Delete(&models.Student{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})

}