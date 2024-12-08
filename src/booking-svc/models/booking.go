package models

import "time"

// Booking defines the structure for a booking API
// Определяет структуру для API бронирования
type Booking struct {
	BookingID int       `gorm:"primaryKey" json:"booking_id"`
	ClientID  int       `json:"client_id"`
	HotelID   int       `json:"hotel_id"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	Price     int       `json:"price"`
	Status    string    `json:"status"`
}
