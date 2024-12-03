package hotels_data

import (
	"github.com/maxturyev/booking-system-project/src/hotel-svc/models"
	"gorm.io/gorm"
)

func GetHotelByPrice(db *gorm.DB, left, right int) []models.Hotel {
	var hotels []models.Hotel
	result := db.Where("rating > ? AND rating < ?", left, right).Find(hotels)
	if result.Error != nil {
		panic("Error")
	}
	return hotels
}
