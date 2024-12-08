package models

type RoomCategory struct {
	RoomCategoryID uint   `gorm:"primaryKey" json:"room_category_id"`
	Name           string `json:"name"`
	Price          string `json:"price"`
	Amount         string `json:"amount"`
	HotelID        uint   `json:"hotel_id"`
}
