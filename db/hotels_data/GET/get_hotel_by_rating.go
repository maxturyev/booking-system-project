package hotels_data

import (
	"github.com/maxturyev/booking-system-project/models"
	"gorm.io/gorm"
)

func GetHotelByRating(db *gorm.DB, left, right int) []models.Hotel {
	if left == 0 {
		left = 0
	}
	if right == 0 {
		right = 5
	}
	var hotels []models.Hotel
	result := db.Where("rating >= ? AND rating <= right", left, right).Find(hotels)
	if result.Error != nil {
		panic("Error")
	}
	return hotels
}
