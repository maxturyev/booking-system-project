package handlers

import "time"

type Hotelier struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	HotelsId  []int  `json:"hotels_id"`
	CreatedOn string `json:"-"`
}

type Hoteliers []*Hotelier

func GetHoteliers() Hoteliers {
	return hotelierList
}

func AddHoteier(h *Hotelier) {
	if len(hotelList) == 0 {
		h.ID = 1
	} else {
		h.ID = hotelierList[len(hotelierList)-1].ID + 1
	}
	hotelierList = append(hotelierList, h)
}

var hotelierList = Hoteliers{
	{
		ID:        1,
		FirstName: "Max",
		LastName:  "Tyrev",
		Login:     "login",
		Password:  "password",
		HotelsId:  make([]int, 0),
		CreatedOn: time.Now().UTC().String(),
	},
}
