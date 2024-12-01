package hotels_data

import (
	"github.com/maxturyev/booking-system-project/models"
	"gorm.io/gorm"
)

func GetHotels(db *gorm.DB) []models.Hotel {
	var hotel []models.Hotel
	result := db.Find(&hotel)
	if result.Error != nil {
		panic("Error")
	}
	return hotel
}
