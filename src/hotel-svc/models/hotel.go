package models

type Hotel struct {
	HotelID        uint   `gorm:"primaryKey;autoIncrement" json:"hotel_id"`
	Name           string `gorm:"index" json:"name"`
	Rating         int    `json:"rating"`
	Country        string `json:"country"`
	Description    string `json:"description"`
	RoomsAvailable int    `json:"rooms_available"`
	Price          string `json:"price"`
	Address        string `gorm:"uniqueIndex" json:"address"`
	HotelierID     uint   `json:"hotelier_id"`
}
