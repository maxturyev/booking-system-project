package models

// Hotelier defines the structure for a hotelier API
type Hotelier struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     string `gorm:"uniqueIndex" json:"login"`
	Password  string `json:"password"`
	HotelsId  int    `json:"hotels_id"`
	CreatedOn string `json:"-"`
}
