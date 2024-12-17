package databases

import (
	"errors"
	"github.com/maxturyev/booking-system-project/booking-svc/models"
	"gorm.io/gorm"
	"log"
)

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
