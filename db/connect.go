package db

import (
	"fmt"

	"github.com/maxturyev/booking-system-project/consts"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewConnection establishes a connection to both databases
func NewConnection() *gorm.DB {
	hotel_db, err := NewHotelConnection()
	if err != nil {

	}

	// NewBookingConnection()

	return hotel_db
}

// NewHotelConnection establishes a connection to hotels database
func NewHotelConnection() (*gorm.DB, error) {
	// Initialize connection to Hotel database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", consts.Host, consts.Port, consts.User, consts.Password, consts.DBHotel)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
