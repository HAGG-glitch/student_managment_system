package controllers

import (
	"net/http"
	"github.com/HAGG-glitch/student_managment_system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateExam(c *gin.Context, db *gorm.DB) {
	var exam models.Exam

	// Validate input
	if err := c.ShouldBindJSON(&exam); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// attempt to create exam in the databse
	if err := db.Create(&exam).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	


	// successfully created
	c.IndentedJSON(http.StatusCreated, exam)
}

func GetExams(c *gin.Context, db *gorm.DB) {
	var exams []models.Exam

	if err := db.Preload("Class").Preload("Student").Preload("Teacher").Find(&exams).Error; err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(exams) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No exams found"})
		return
	}

	c.IndentedJSON(http.StatusOK, exams)
}


// Find by ID
func GetExamByID(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	if err := db.Find(&models.Exam{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Exam not found"} )

}



// UpdateStudent updates a student's information by ID
func UpdateExam(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	var exam models.Exam
	// Validate input
	if err := c.ShouldBindJSON(&exam); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// Attempt to update student in the database
	if err := db.Model(&models.Exam{}).Where("id = ?", id).Updates(exam).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Successfully updated
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Exam updated successfully"})
}



// DeleteStudent deletes a student by ID
func DeleteExam(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	if err := db.Delete(&models.Exam{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Exam is deleted successfully"})

}