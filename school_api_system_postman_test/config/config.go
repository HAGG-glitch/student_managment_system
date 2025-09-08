package config

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error, bool) {
	// Load environment variables
	host := "localhost"
	port := "5432"
	user := "super_user"
	password := "123456"
	dbname := "school_system_2"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err, false
	}

	log.Println("Database connection established")
	return db, nil, true
}