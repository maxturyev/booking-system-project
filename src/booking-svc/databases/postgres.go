// this file not have a all checking ERRORS in the future will be fixed today 08.12.2024
package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/maxturyev/booking-system-project/booking-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Establish a connection to booking database
func Init() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize connection to Booking database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("Host"),
		os.Getenv("Port"), os.Getenv("User"), os.Getenv("Password"), os.Getenv("BookingDB"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	// Migrates models to the database as tables
	migrateDB(db)

	return db, nil
}

// Migration model(struct) to database(Postgres)
func migrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.Client{}, &models.Booking{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	createConstraints(db)
}

// createConstraints is used to create a foreign key constraint
func createConstraints(db *gorm.DB) {
	if err := db.Migrator().CreateConstraint(&models.Booking{}, "Booking"); err != nil {
		log.Printf("Failed to create constraint: %v", err)
	}
}
