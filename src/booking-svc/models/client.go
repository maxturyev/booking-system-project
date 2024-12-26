package models

// Client defines the structure for a client API
type Client struct {
	ClientID  int      `gorm:"primaryKey;autoIncrement" json:"client_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `gorm:"uniqueIndex" json:"email"`
	Phone     string   `gorm:"uniqueIndex" json:"phone"`
	Login     string   `gorm:"uniqueIndex" json:"login"`
	Password  string   `json:"-"`
	Bookings  Bookings `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

// Clients is a collection of Client
type Clients []*Client
