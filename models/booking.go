package models

import "time"

// Booking defines the structure for a booking API
// Определяет структуру для API бронирования
type Booking struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	ClientID     int       `json:"client_id"`
	HotelID      int       `json:"hotel_id"`
	DateStart    time.Time `json:"date_start"`
	DateEnd      time.Time `json:"date_end"`
	RoomCategory string    `json:"room_category"`
	RoomCount    int       `json:"room_count"`
	Price        int       `json:"price"`
	GuestAmount  int       `json:"guest_amount"`
	IsCancelled  bool      `json:"is_cancelled"`
}
