package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/maxturyev/booking-system-project/booking-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to booking database
func ConnectDB() *gorm.DB {
	// Initialize connection to Booking database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("HOST"),
		os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASS"), os.Getenv("BOOKING_DB"))
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
