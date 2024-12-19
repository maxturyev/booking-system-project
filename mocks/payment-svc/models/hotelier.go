package models

type Hotelier struct {
	HotelierID uint    `gorm:"primaryKey" json:"hotelier_id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Email      string  `gorm:"uniqueIndex" json:"email"`
	Phone      string  `gorm:"uniqueIndex" json:"phone"`
	Login      string  `gorm:"uniqueIndex" json:"login"`
	Password   string  `json:"-"`
	Hotels     []Hotel `gorm:"foreignKey:HotelierID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"hotels"`
}
