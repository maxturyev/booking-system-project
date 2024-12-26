package models

// Hotelier defines the structure for a hotelier API
type Hotelier struct {
	HotelierID uint   `gorm:"primaryKey;autoIncrement" json:"hotelier_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `gorm:"uniqueIndex" json:"email"`
	Phone      string `gorm:"uniqueIndex" json:"phone"`
	Login      string `gorm:"uniqueIndex" json:"login"`
	Password   string `json:"-"`
	Hotels     Hotels `gorm:"foreignKey:HotelierID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

// Hoteliers is a collection of Hotelier
type Hoteliers []*Hotelier
