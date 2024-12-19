package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/maxturyev/booking-system-project/payment-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to hotels database
func ConnectDB() *gorm.DB {
	// Load postgres server config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize connection to Hotels database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("Host"),
		os.Getenv("Port"), os.Getenv("User"), os.Getenv("Password"), os.Getenv("HotelDB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate models to the database as tables
	if err := db.AutoMigrate(&models.Hotelier{}, &models.Hotel{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return db
}
