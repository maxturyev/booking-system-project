package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/maxturyev/booking-system-project/src/booking-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AddBookings(db *gorm.DB, bookings models.Bookings) {
	for _, booking := range bookings {
		db.Create(&booking)
	}
}

func AddClients(db *gorm.DB, clients models.Clients) {
	for _, client := range clients {
		client.Email = fmt.Sprintf("%s@%s.com", client.FirstName, client.LastName)
		client.Password = fmt.Sprintf("%s_%s", client.FirstName, client.LastName)
		db.Create(&client)
	}
}

// ConnectDB establishes a connection to Booking database
func ConnectDB() *gorm.DB {
	// Initialize connection to Bookings database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("BOOKING_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate models to the database as tables
	if err := db.AutoMigrate(&models.Client{}, &models.Booking{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	bookings := models.Bookings{
		{HotelID: 1, DateStart: time.Now(), DateEnd: time.Now().AddDate(0, 0, 7), Price: 150.0, Status: "confirmed", ClientID: 1},
		{HotelID: 2, DateStart: time.Now().AddDate(0, 0, 1), DateEnd: time.Now().AddDate(0, 0, 8), Price: 200.0, Status: "pending", ClientID: 2},
	}

	clients := models.Clients{
		{FirstName: "Alice", LastName: "Smith", Phone: "1234567890", Login: "alicesmith"},
		{FirstName: "Bob", LastName: "Johnson", Phone: "0987654321", Login: "bobjohnson"},
	}

	// Add clients first to satisfy foreign key constraint
	AddClients(db, clients)
	AddBookings(db, bookings)

	fmt.Println("Sample data added successfully")

	return db
}
