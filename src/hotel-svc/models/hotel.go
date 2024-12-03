package models

// Hotel defines the structure for a hotel API
// Определяет структуру для API отеля
type Hotel struct {
	HotelID        int    `gorm:"primaryKey" json:"hotel_id"`
	Name           string `json:"name"`
	HotelierID     int    `json:"hotelier_id"`
	Rating         int    `json:"rating"`
	Country        string `json:"country"`
	Address        string `gorm:"uniqueIndex" json:"address"`
	RoomsAvailable int    `json:"rooms_available"`
}
