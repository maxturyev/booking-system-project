package models

// Hotel defines the structure for a hotel API
type Hotel struct {
	ID             int    `gorm:"primaryKey" json:"id"`
	Name           string `json:"name"`
	HotelierID     int    `json:"hotelier_id"`
	Rating         int    `json:"rating"`
	Country        string `json:"country"`
	Address        string `gorm:"uniqueIndex" json:"address"`
	RoomsAvailable int    `json:"rooms_available"`
	CreatedOn      string `json:"-"`
	UpdatedOn      string `json:"-"`
	DeletedOn      string `json:"-"`
}
