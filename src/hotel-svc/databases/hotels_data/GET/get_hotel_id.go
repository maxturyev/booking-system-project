package hotels_data

import (
	"github.com/maxturyev/booking-system-project/src/hotel-svc/models"
	"gorm.io/gorm"
)

// GetHotelByID adds a new hotel to the data store
func GetHotelByID(db *gorm.DB, id int) models.Hotel {
	var hotel models.Hotel
	result := db.First(&hotel, id)
	if result.Error != nil {
		panic("Error")
	}
	return hotel
}
