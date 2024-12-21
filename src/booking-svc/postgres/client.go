package postgres

import (
	"errors"
	"github.com/maxturyev/booking-system-project/src/booking-svc/models"
	"gorm.io/gorm"
	"log"
)

// CreateClient creates a client to the database
func CreateClient(db *gorm.DB, client models.Client) {
	db.Create(&client)
}

// SelectClients returns all clients from the database
func SelectClients(db *gorm.DB) models.Clients {
	var clients models.Clients

	result := db.Find(&clients)
	if result.Error != nil {
		panic("Error")
	}

	return clients
}

// UpdateClient updates client info in the database
func UpdateClient(db *gorm.DB, client models.Client) error {
	log.Println("entered postgres update")
	var existing models.Client

	result := db.First(&existing, client.ClientID)
	if result.Error != nil {
		return result.Error
	}

	result = db.Model(&existing).Updates(client)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SelectClientByID(db *gorm.DB, id int) (models.Client, error) {
	var client models.Client

	result := db.First(&client, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("Запись не найдена")
	}

	return client, result.Error
}
