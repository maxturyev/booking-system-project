package models

// Client defines the structure for a client API
// Определяет структуру для API клиента
type Client struct {
	ClientID  int    `gorm:"primaryKey" json:"client_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `gorm:"uniqueIndex" json:"email"`
	Country   string `gorm:"uniqueIndex" json:"country"`
	Phone     string `gorm:"uniqueIndex" json:"phone"`
	Login     string `gorm:"uniqueIndex" json:"login"`
	Password  string `json:"password"`
}
