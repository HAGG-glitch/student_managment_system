package controllers

import "github.com/gin-gonic/gin"
import "net/http"

// StudentDashboard handles the student dashboard
func StudentDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Student Dashboard",
	})
}

// TeacherDashboard handles the teacher dashboard
func TeacherDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Teacher Dashboard",
	})
}

// AdminDashboard handles the admin dashboard
func AdminDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Admin Dashboard",
	})
}
