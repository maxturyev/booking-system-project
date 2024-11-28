package handlers

// Client defines the structure for an API client
type Client struct {
	ID        int    `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"hotelier"`
	Email     string `json:"rating"`
	Country   string `json:"country"`
	Phone     string `json:"phone"`
	Bookings  []int  `json:"bookings"`
}

// // NewClient creates a client
// func NewClient(id int, fname string, lname string, email string, country string, phone string) *Client {
// 	return &Client{id, fname, lname, email, country, phone}
// }

// Clients is a collection of Client
type Clients []*Client
