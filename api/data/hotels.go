package data

import (
	"time"
)

// Hotel defines the structure for an API hotel
type Hotel struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Hotelier  int    `json:"hotelier"`
	Rating    int    `json:"rating"`
	Country   string `json:"country"`
	Address   string `json:"address"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Hotels is a collection of Product
type Hotels []*Hotel

// GetHotels returns a list of hotels
func GetHotels() Hotels {
	return hotelList
}

func GetHotel(h Hotel) Hotel {
	return h
}

// AddHotel adds a new hotel to the data store
func AddHotel(h *Hotel) {
	h.ID = hotelList[len(hotelList)].ID
	hotelList = append(hotelList, h)
}

// Example data store
var hotelList = Hotels{
	{
		ID:        1,
		Name:      "Hilton Moscow Leningradskaya",
		Hotelier:  1,
		Rating:    5,
		Country:   "Russia",
		Address:   "Kalanchevskaya street 21/40",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        2,
		Name:      "Cavalieri Hotel",
		Hotelier:  1,
		Rating:    4,
		Country:   "Greece",
		Address:   "Kapodistriou street 4",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}