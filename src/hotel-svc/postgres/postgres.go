package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/maxturyev/booking-system-project/src/hotel-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AddHotels(db *gorm.DB, hotels models.Hotels) {
	for _, hotel := range hotels {
		db.Create(&hotel)
	}
}

func AddHoteliers(db *gorm.DB, hoteliers models.Hoteliers) {
	for _, hotelier := range hoteliers {
		hotelier.Email = fmt.Sprintf("%s@%s.com", hotelier.FirstName, hotelier.LastName)
		hotelier.Password = fmt.Sprintf("%s_%s", hotelier.FirstName, hotelier.LastName)
		db.Create(&hotelier)
	}
}

// ConnectDB establishes a connection to Hotels database
func ConnectDB() *gorm.DB {
	// Initialize connection to Hotels database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("HOTEL_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate models to the database as tables
	if err := db.AutoMigrate(&models.Hotelier{}, &models.Hotel{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	hotels := models.Hotels{
		{Name: "Hotel A", Rating: 4, Country: "USA", Description: "Luxury hotel", RoomsAvailable: 100, RoomPrice: 150.0, Address: "123 Main St", HotelierID: 1},
		{Name: "Hotel B", Rating: 5, Country: "Canada", Description: "Boutique hotel", RoomsAvailable: 50, RoomPrice: 200.0, Address: "456 Elm St", HotelierID: 2},
	}

	hoteliers := models.Hoteliers{
		{FirstName: "John", LastName: "Doe", Phone: "1234567890", Login: "johndoe"},
		{FirstName: "Jane", LastName: "Smith", Phone: "0987654321", Login: "janesmith"},
	}

	// Add hoteliers first to satisfy foreign key constraint
	AddHoteliers(db, hoteliers)
	AddHotels(db, hotels)

	fmt.Println("Sample data added successfully")

	return db
}
