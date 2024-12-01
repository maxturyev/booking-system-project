package models

// Client defines the structure for a client API
type Client struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"hotelier"`
	Email     string `gorm:"uniqueIndex" json:"rating"`
	Country   string `gorm:"uniqueIndex" json:"country"`
	Phone     string `json:"phone"`
	Bookings  int    `json:"bookings"`
}
