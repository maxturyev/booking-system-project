package db

import (
	"github.com/maxturyev/booking-system-project/hotel-svc/models"
	"gorm.io/gorm"
)

// CreateHotelier creates a hotelier to the database
func CreateHotelier(db *gorm.DB, hotelier models.Hotelier) error {
	result := db.Create(&hotelier)
	return result.Error
}

// SelectHoteliers returns all hoteliers from the database
func SelectHoteliers(db *gorm.DB) (models.Hoteliers, error) {
	var hoteliers models.Hoteliers

	result := db.Find(&hoteliers)
	if result.Error != nil {
		return hoteliers, result.Error
	}

	return hoteliers, nil
}
