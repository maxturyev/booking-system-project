package hotels_data

import (
	"github.com/maxturyev/booking-system-project/models"
	"gorm.io/gorm"
)

func GetHotelByRating(db *gorm.DB, stars ...int) []models.Hotel {
	var hotels []models.Hotel

	result := db.Where("rating IN ?", stars).Find(hotels)
	if result.Error != nil {
		panic("Error")
	}

	return hotels
}
