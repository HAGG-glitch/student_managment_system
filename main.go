package main

import (
	"log"

	"github.com/HAGG-glitch/student_managment_system/config"
	"github.com/HAGG-glitch/student_managment_system/models"
	"github.com/HAGG-glitch/student_managment_system/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Connect to DB
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }

    db, err, _ := config.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }


// Independent tables
err = db.AutoMigrate(
    &models.Class{},
    &models.Subject{},
    &models.User{},
)
if err != nil {
    log.Fatalf("Migration failed (independent tables): %v", err)
}

// Dependent tables
err = db.AutoMigrate(
    &models.Teacher{},
    &models.Student{},
    &models.Enrollment{},
    &models.Exam{},
    &models.Attendance{},
)
if err != nil {
    log.Fatalf("Migration failed (dependent tables): %v", err)
} else {
    log.Println("Migration successful")
}

	// Initialize router
	r := gin.Default()

	// Load routes
	routes.RegisterRoutes(r, db)

	// Start server
	r.Run(":9000") // Runs on localhost:8080
}
