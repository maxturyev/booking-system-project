package data

import (
	"time"
)

// Hotel defines the structure for an API hotel
type Hotel struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Hotelier       int    `json:"hotelier"`
	Rating         int    `json:"rating"`
	Country        string `json:"country"`
	Address        string `json:"address"`
	RoomsAvailable int    `json:"rooms_available"`
	CreatedOn      string `json:"-"`
	UpdatedOn      string `json:"-"`
	DeletedOn      string `json:"-"`
}

// Hotels is a collection of Product
type Hotels []*Hotel

// GetHotels returns a list of hotels
func GetHotels() Hotels {
	return hotelList
}

// AddHotel adds a new hotel to the data store
func AddHotel(h *Hotel) {
	h.ID = hotelList[len(hotelList)-1].ID + 1
	hotelList = append(hotelList, h)
}

// Example data store
var hotelList = Hotels{
	{
		ID:             1,
		Name:           "Hilton Moscow Leningradskaya",
		Hotelier:       1,
		Rating:         5,
		Country:        "Russia",
		Address:        "Kalanchevskaya street 21/40",
		RoomsAvailable: 123,
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
	{
		ID:             2,
		Name:           "Cavalieri Hotel",
		Hotelier:       1,
		Rating:         4,
		Country:        "Greece",
		Address:        "Kapodistriou street 4",
		RoomsAvailable: 0,
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
}
