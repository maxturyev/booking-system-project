package databases

import (
	"github.com/maxturyev/booking-system-project/booking-svc/models"
	"gorm.io/gorm"
)

// AddClient adds a client to the database
func AddClient(db *gorm.DB, client models.Client) {
	db.Create(&client)
}

// GetClients gets all clients in the database
func GetClients(db *gorm.DB) models.Clients {
	var clients models.Clients
	result := db.Find(&clients)
	if result.Error != nil {
		panic("Error")
	}
	return clients
}
