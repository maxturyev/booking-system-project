package data

import (
	"time"
)

// Booking defines the structure for an API hotel
type Booking struct {
	ID           int    `json:"id"`
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

// Bookings is a collection of Booking
type Bookings []*Booking

// GetBookings returns a list of bookings
func GetBookings() Bookings {
	return bookingList
}

// CreateBooking c>reates a new booking
func CreateBooking(b *Booking) {
	b.ID = bookingList[len(bookingList)-1].ID + 1
	bookingList = append(bookingList, b)
}

// Example data store
var bookingList = Bookings{
	{
		ID:           1,
		ClientID:     1,
		HotelID:      1,
		RoomCategory: "double",
		RoomCount:    1,
		Price:        100,
		GuestCount:   2,
		CreatedOn:    time.Now().UTC().String(),
		UpdatedOn:    time.Now().UTC().String(),
	},
	{
		ID:           2,
		ClientID:     1,
		HotelID:      1,
		RoomCategory: "single",
		RoomCount:    1,
		Price:        80,
		GuestCount:   1,
		CreatedOn:    time.Now().UTC().String(),
		UpdatedOn:    time.Now().UTC().String(),
	},
	{
		ID:           3,
		ClientID:     1,
		HotelID:      2,
		RoomCategory: "suite",
		RoomCount:    1,
		Price:        120,
		GuestCount:   2,
		CreatedOn:    time.Now().UTC().String(),
		UpdatedOn:    time.Now().UTC().String(),
	},
	{
		ID:           4,
		ClientID:     1,
		HotelID:      2,
		RoomCategory: "studio",
		RoomCount:    2,
		Price:        100,
		GuestCount:   2,
		CreatedOn:    time.Now().UTC().String(),
		UpdatedOn:    time.Now().UTC().String(),
	},
}
