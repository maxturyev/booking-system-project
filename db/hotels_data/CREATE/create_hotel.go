package hotels_data

import (
	"github.com/maxturyev/booking-system-project/models"
	"gorm.io/gorm"
)

// CreateHotel adds a new hotel to the database
func CreateHotel(db *gorm.DB, hotel models.Hotel) {
	db.Create(&hotel)
}
