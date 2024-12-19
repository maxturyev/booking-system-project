package models

// Hotel defines the structure for a hotel API
type Hotel struct {
	HotelID        uint    `gorm:"primaryKey;autoIncrement" json:"hotel_id"`
	Name           string  `gorm:"index" json:"name"`
	Rating         int     `json:"rating"`
	Country        string  `json:"country"`
	Description    string  `json:"description"`
	RoomsAvailable int     `json:"rooms_available"`
	RoomPrice      float32 `json:"room_price"`
	Address        string  `gorm:"uniqueIndex" json:"address"`
	HotelierID     uint    `json:"hotelier_id"`
}

// Hotels is a collection of Hotel
type Hotels []*Hotel
