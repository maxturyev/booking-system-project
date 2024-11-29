package models

// Client defines the structure for a client API
type Client struct {
	ID        int    `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"hotelier"`
	Email     string `json:"rating"`
	Country   string `json:"country"`
	Phone     string `json:"phone"`
	Bookings  []int  `json:"bookings"`
}
