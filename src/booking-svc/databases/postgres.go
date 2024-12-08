// this file not have a all checking ERRORS in the future will be fixed today 08.12.2024
package databases

import (
	"errors"
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

// this function need for a creating foreign key in database because GORM can not automigration with slices
func createConstraints(db *gorm.DB) {
	if err := db.Migrator().CreateConstraint(&models.Booking{}, "Booking"); err != nil {
		log.Printf("Failed to create constraint: %v", err)
	}
}

// Update hotel with need field
func UpdateBooking(db *gorm.DB, booking models.Booking) error {
	var existing models.Booking
	result := db.First(&existing, booking.BookingID)

	// Check error
	if result.Error != nil {
		return result.Error
	}

	result = db.Model(&existing).Updates(booking)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// function which work with booking handler
// get all bookings in system
func GetBookings(db *gorm.DB) []models.Booking {
	var booking []models.Booking
	result := db.Find(&booking)
	if result.Error != nil {
		panic("Error")
	}
	return booking
}

// function which work with booking handler
// get booking from database by id
func GetBookingByID(db *gorm.DB, id int) models.Booking {
	var hotel models.Booking
	result := db.First(&hotel, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("Запись не найдена")
	}
	return hotel
}

func DeleteBookingByID(db *gorm.DB, id int) bool {
	result := db.Delete(&models.Booking{}, id)
	if result.Error != nil {
		panic("Error")
	}
	return true
}

// creating bookings
// function can use only client
func CreateBooking(db *gorm.DB, booking models.Booking) {
	db.Create(&booking)
}
