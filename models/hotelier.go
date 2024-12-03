package models

// Hotelier defines the structure for a hotelier API
// Определяет структуру для API владельца
type Hotelier struct {
	HotelierID int    `gorm:"primaryKey" json:"hotelier_id"`
	HotelID    int    `gorm:"uniqueIndex" json:"hotel_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `gorm:"uniqueIndex" json:"email"`
	Phone      string `gorm:"uniqueIndex" json:"phone"`
	Login      string `gorm:"uniqueIndex" json:"login"`
	Password   string `json:"password"`
	HotelsId   int    `json:"hotels_id"`
}
