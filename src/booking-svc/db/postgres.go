package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/maxturyev/booking-system-project/booking-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to booking database
func ConnectDB() *gorm.DB {
	// Load postgres server config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize connection to Hotels database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("Host"),
		os.Getenv("Port"), os.Getenv("User"), os.Getenv("Password"), os.Getenv("BookingDB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrates models to the database as tables
	if err := db.AutoMigrate(&models.Client{}, &models.Booking{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return db
}
