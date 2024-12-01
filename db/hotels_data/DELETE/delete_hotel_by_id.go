package hotels_data

import (
	hotels_data "github.com/maxturyev/booking-system-project/db/hotels_data/GET"
	"github.com/maxturyev/booking-system-project/models"
	"gorm.io/gorm"
)

func DeleteHotelByID(db *gorm.DB, id int) bool {
	hotel := hotels_data.GetHotelByID(db, id)
	nullHotel := models.Hotel{}
	if hotel == nullHotel {
		return false
	}
	result := db.Delete(&models.Hotel{}, id)
	if result.Error != nil {
		panic("Error")
	}
	return true
}
