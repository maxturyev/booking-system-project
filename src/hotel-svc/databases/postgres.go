package databases

import (
	"fmt"

	"github.com/maxturyev/booking-system-project/src/hotel-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Host      = "localhost"
	Port      = "5433"
	User      = "postgres"
	Password  = "Alan2805"
	DBBooking = "BookingData"
	DBHotel   = "postgres"
)

// NewHotelConnection establishes a connection to hotels database
func Init() (*gorm.DB, error) {
	// Initialize connection to Hotel database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", Host, Port, User, Password, DBHotel)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	// Migrates models to the database as tables
	db.AutoMigrate(&models.Hotel{})

	return db, nil
}
