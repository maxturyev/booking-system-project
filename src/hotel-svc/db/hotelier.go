package db

import (
	"github.com/maxturyev/booking-system-project/hotel-svc/models"
	"gorm.io/gorm"
)

// CreateHotelier creates a hotelier to the database
func CreateHotelier(db *gorm.DB, hotelier models.Hotelier) {
	db.Create(&hotelier)
}

// SelectHoteliers returns all hoteliers from the database
func SelectHoteliers(db *gorm.DB) models.Hoteliers {
	var hoteliers models.Hoteliers

	result := db.Find(&hoteliers)
	if result.Error != nil {
		panic("Error")
	}

	return hoteliers
}
