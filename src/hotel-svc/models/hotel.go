package models

type Hotel struct {
	HotelID        uint           `gorm:"primaryKey;autoIncrement" json:"hotel_id"`
	Name           string         `gorm:"index" json:"name"`
	Rating         int            `json:"rating"`
	Country        string         `json:"country"`
	Description    string         `json:"description"`
	Address        string         `gorm:"uniqueIndex" json:"address"`
	HotelierID     uint           `json:"hotelier_id"`
	RoomCategories []RoomCategory `gorm:"foreignKey:HotelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room_categories"`
}
