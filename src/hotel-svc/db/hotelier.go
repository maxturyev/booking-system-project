package db

import (
	"github.com/maxturyev/booking-system-project/hotel-svc/models"
	"gorm.io/gorm"
)

// CreateHotelier adds a hotelier to the database
func CreateHotelier(db *gorm.DB, hotelier models.Hotelier) {
	db.Create(&hotelier)
}

// GetHoteliers fetches hoteliers from the database
func GetHoteliers(db *gorm.DB) []models.Hotelier {
	var hotelier []models.Hotelier
	result := db.Find(&hotelier)
	if result.Error != nil {
		panic("Error")
	}
	return hotelier
}
