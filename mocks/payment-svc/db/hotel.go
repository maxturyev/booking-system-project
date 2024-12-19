package db

import (
	"errors"
	"log"

	"github.com/maxturyev/booking-system-project/payment-svc/models"
	"gorm.io/gorm"
)

// CreateHotel creates a hotel to the database
func CreateHotel(db *gorm.DB, hotel models.Hotel) {
	db.Create(&hotel)
}

// SelectHotels returns all hotels from the database
func SelectHotels(db *gorm.DB) []models.Hotel {
	var hotel []models.Hotel

	result := db.Find(&hotel)
	if result.Error != nil {
		panic("Error")
	}

	return hotel
}

// UpdateHotel updates hotel info in the database
func UpdateHotel(db *gorm.DB, hotel models.Hotel) error {
	var existing models.Hotel

	result := db.First(&existing, hotel.HotelID)
	if result.Error != nil {
		return result.Error
	}

	result = db.Model(&existing).Updates(hotel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// SelectHotelByID returns a hotel from the database by ID
func SelectHotelByID(db *gorm.DB, id int) models.Hotel {
	var hotel models.Hotel

	result := db.First(&hotel, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("Запись не найдена")
	}

	return hotel
}

// GetHotelByRating fetches hotels from the database by rating
func GetHotelByRating(db *gorm.DB, stars ...int) []models.Hotel {
	var hotels []models.Hotel

	result := db.Where("rating IN ?", stars).Find(hotels)
	if result.Error != nil {
		panic("Error")
	}

	return hotels
}

// GetHotelByPrice fetches hotels from the database by price
func GetHotelByPrice(db *gorm.DB, left, right int) []models.Hotel {
	var hotels []models.Hotel
	result := db.Where("rating > ? AND rating < ?", left, right).Find(hotels)
	if result.Error != nil {
		panic("Error")
	}
	return hotels
}

// DeleteHotelByID removes a hotel from the database
func DeleteHotelByID(db *gorm.DB, id int) bool {
	result := db.Delete(&models.Hotel{}, id)
	if result.Error != nil {
		panic("Error")
	}
	return true
}
