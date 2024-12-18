package db

import (
	"errors"
	"github.com/maxturyev/booking-system-project/hotel-svc/models"
	"gorm.io/gorm"
	"log"
)

// CreateHotel adds a hotel to the database
func CreateHotel(db *gorm.DB, hotel models.Hotel) {
	db.Create(&hotel)
}

// UpdateHotel updates hotel info
func UpdateHotel(db *gorm.DB, hotel models.Hotel) error {
	var existing models.Hotel
	result := db.First(&existing, hotel.HotelID)
	// Check error
	if result.Error != nil {
		return result.Error
	}

	result = db.Model(&existing).Updates(hotel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetHotels fetches hotels from the database
func GetHotels(db *gorm.DB) []models.Hotel {
	var hotel []models.Hotel
	result := db.Find(&hotel)
	if result.Error != nil {
		panic("Error")
	}
	return hotel
}

// GetHotelByID fetches hotels from the database by ID
func GetHotelByID(db *gorm.DB, id int) models.Hotel {
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
