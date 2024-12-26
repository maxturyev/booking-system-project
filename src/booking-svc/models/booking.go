package models

import "time"

// Booking defines the structure for a booking API
type Booking struct {
	BookingID int       `gorm:"primaryKey" json:"booking_id"`
	HotelID   int       `json:"hotel_id"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	Price     float32   `json:"price"`
	Status    string    `json:"status"`
	ClientID  int       `json:"client_id"`
}

// Bookings is a collection of Booking
type Bookings []*Booking
