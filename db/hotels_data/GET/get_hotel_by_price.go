package hotels_data

import (
	"github.com/maxturyev/booking-system-project/consts"
	"github.com/maxturyev/booking-system-project/models"
	"gorm.io/gorm"
)

func GetHotelByPrice(db *gorm.DB, left, right int) []models.Hotel {
	if right == 0 {
		right = consts.MaxInt
	}
	var hotels []models.Hotel
	result := db.Where("rating > ? AND rating < ?", left, right).Find(hotels)
	if result.Error != nil {
		panic("Error")
	}
	return hotels
}
