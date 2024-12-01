package models

// Booking defines the structure for an booking API
type Booking struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	ClientID     int    `json:"client_id"`
	HotelID      int    `json:"hotel_id"`
	RoomCategory string `json:"room_category"`
	RoomCount    int    `json:"room_count"`
	Price        int    `json:"price"`
	GuestCount   int    `json:"guest_count"`
	CreatedOn    string `json:"-"`
	UpdatedOn    string `json:"-"`
	DeletedOn    string `json:"-"`
}
