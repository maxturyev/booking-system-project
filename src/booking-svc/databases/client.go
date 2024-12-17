package databases

import (
	"github.com/maxturyev/booking-system-project/booking-svc/models"
	"gorm.io/gorm"
	"log"
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

// Update hotel with need field
func UpdateClient(db *gorm.DB, client models.Client) error {
	log.Println("entered db update")
	var existing models.Client

	result := db.First(&existing, client.ClientID)

	// Check error
	if result.Error != nil {
		return result.Error
	}

	result = db.Model(&existing).Updates(client)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
