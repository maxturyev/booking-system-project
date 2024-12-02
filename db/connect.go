package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/maxturyev/booking-system-project/consts"
	"github.com/maxturyev/booking-system-project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

// Initialize databases
// Creating tables from models
func Init() *gorm.DB {
	hotel_db, err := NewHotelConnection()
	if err != nil {
		panic("ERROR")
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

	// Migrates models to the database as tables
	db.AutoMigrate(&models.Booking{}, &models.Client{}, &models.Hotel{}, &models.Hotelier{})

	return db, nil
}
