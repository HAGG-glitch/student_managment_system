package routes

import (
	"github.com/HAGG-glitch/student_managment_system/controllers"
	"github.com/HAGG-glitch/student_managment_system/middleware"

	 sec  "github.com/HAGG-glitch/student_managment_system/security"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes sets up the routes for the application

func RegisterRoutes(r *gin.Engine, db *gorm.DB){
	api := r.Group("/api/v1")
	{
		// =========================
		// Student routes
		// =========================
		api.GET("/students", func(c *gin.Context) {controllers.GetStudents(c, db)})
		api.GET("/students/:id", func(c *gin.Context) {controllers.GetStudentByID(c, db)})
		api.POST("/students", func(c *gin.Context) {controllers.CreateStudent(c, db)})
		api.PUT("/students/:id", func(c *gin.Context) {controllers.UpdateStudent(c, db)})
		api.DELETE("/students/:id", func(c *gin.Context) {controllers.DeleteStudent(c, db)})
		
		// =========================
		// Teacher routes
		// =========================

		api.GET("/teachers", func(c *gin.Context) {controllers.GetTeachers(c, db)})
		api.GET("/teachers/:id", func(c *gin.Context) {controllers.GetTeacherByID(c, db)})
		api.POST("/teachers", func(c *gin.Context) {controllers.CreateTeacher(c, db)})
		api.PUT("/teachers/:id", func(c *gin.Context) {controllers.UpdateTeacher(c, db)})
		api.DELETE("/teachers/:id", func(c *gin.Context) {controllers.DeleteTeacher(c, db)})
		
		// =========================
		// Class routes
		// =========================
		api.GET("/classes", func(c *gin.Context) {controllers.GetClasses(c, db)})
		api.GET("/classes/:id", func(c *gin.Context) {controllers.GetClassByID(c, db)})
		api.POST("/classes", func(c *gin.Context) {controllers.CreateClass(c, db)})
		api.PUT("/classes/:id", func(c *gin.Context) {controllers.UpdateClass(c, db)})
		api.DELETE("/classes/:id", func(c *gin.Context) {controllers.DeleteClass(c, db)})

		// =========================
		//Subject routes
		// =========================

		api.GET("/subjects", func(c *gin.Context) {controllers.GetSubjects(c, db)})
		api.GET("/subjects/:id", func(c *gin.Context) {controllers.GetSubjectByID(c, db)})
		api.POST("/subjects", func(c *gin.Context) {controllers.CreateSubject(c, db)})
		api.PUT("/subjects/:id", func(c *gin.Context) {controllers.UpdateSubject(c, db)})
		api.DELETE("/subjects/:id", func(c *gin.Context) {controllers.DeleteSubject(c, db)})

		// =========================
		// User routes
		// =========================

		api.POST("/register", func(c *gin.Context) { controllers.RegisterUser(c, db) })
		api.POST("/login", func(c *gin.Context) { controllers.LoginUser(c, db) })
		
		// =========================
		// Dashboard routes
		// =========================
		api.GET("/student-dashboard", middleware.AuthMiddleware(), sec.RoleMiddleware("student", "teacher", "admin"), controllers.StudentDashboard)
		api.GET("/teacher-dashboard", middleware.AuthMiddleware(), sec.RoleMiddleware("teacher", "admin"), controllers.TeacherDashboard)
		api.GET("/admin-dashboard", middleware.AuthMiddleware(), sec.RoleMiddleware("admin"), controllers.AdminDashboard)

		// =========================
		// Profile
		// =========================
		api.GET("/profile", middleware.AuthMiddleware(), func(c *gin.Context) {controllers.GetProfile(c, db)})
		
		// =========================
		//attendance routes
		// =========================
		api.POST("/attendance", func(c *gin.Context) {controllers.MarkAttendance(c, db)})
		api.GET("/attendance/student/:student_id", func(c *gin.Context) {controllers.GetAttendanceByStudent(c, db)})
		api.GET("/attendance/class/:class_id", func(c *gin.Context) {controllers.GetAttendanceByClass(c, db)})
		
		// =========================
		//Exam routes
		// =========================
		api.POST("/exams", func(c *gin.Context) {controllers.CreateExam(c, db)})
		api.GET("/exams", func(c *gin.Context) {controllers.GetExams(c, db)})
		api.GET("/exams/:id", func(c *gin.Context) {controllers.GetExamByID(c, db)})
		api.PUT("/exams/:id", func(c *gin.Context) {controllers.UpdateExam(c, db)})
		api.DELETE("/exams/:id", func(c *gin.Context) {controllers.DeleteExam(c, db)})

		// =========================
		// Enrollment routes
		// =========================
		api.POST("/enrollments", func(c *gin.Context) {controllers.CreateEnrollment(c, db)})
		api.GET("/enrollments", func(c *gin.Context) {controllers.GetEnrollments(c, db)})
		api.GET("/enrollments/:id", func(c *gin.Context) {controllers.GetEnrollmentByID(c, db)})
		api.PUT("/enrollments/:id", func(c *gin.Context) {controllers.UpdateEnrollment(c, db)})
		api.DELETE("/enrollments/:id", func(c *gin.Context) {controllers.DeleteEnrollment(c, db)})


	}
}
