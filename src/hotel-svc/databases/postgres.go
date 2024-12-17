// this file not have a all checking ERRORS in the future will be fixed today 08.12.2024
package databases

import (
	"errors"
	"fmt"
	"log"

	"github.com/maxturyev/booking-system-project/hotel-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Host      = "localhost"
	Port      = "5433"
	User      = "postgres"
	Password  = "Alan2805"
	DBBooking = "BookingData"
	DBHotel   = "postgres"
)

// NewHotelConnection establishes a connection to hotels database
func Init() (*gorm.DB, error) {
	// Initialize connection to Hotel database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", Host, Port, User, Password, DBHotel)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	migrateDB(db)
	// Migrates models to the database as tables
	//db.AutoMigrate(&models.RoomCategory{}, &models.Hotel{}, &models.Hotelier{})
	//db.AutoMigrate(&models.Hotel{}, &models.RoomCategory{}, &models.Hotelier{})

	if err != nil {
		log.Fatalf("Ошибка миграции моделей")
	}
	return db, nil
}

// Migration model(struct) to database(Postgres)
func migrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.Hotelier{}, &models.Hotel{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	createConstraints(db)
}

// this function need for a creating foreing key in database because GORM can not automigration with slices
func createConstraints(db *gorm.DB) {
	if err := db.Migrator().CreateConstraint(&models.Hotel{}, "Hotelier"); err != nil {
		log.Printf("Failed to create constraint: %v", err)
	}
	if err := db.Migrator().CreateConstraint(&models.Hotel{}, "fk_hotels_hoteliers"); err != nil {
		log.Printf("Failed to create foreign key constraint: %v", err)
	}
}

// function which work with hotel handler
// Update hotel with need field
func UpdateHotel(db *gorm.DB, hotel models.Hotel) error {
	var existing models.Hotel
	result := db.First(&existing, hotel.HotelID)
	// Check error
	if result.Error != nil {
		return result.Error
	}

	result = db.Model(&existing).Updates(hotel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetHoteliers(db *gorm.DB) []models.Hotelier {
	var hotelier []models.Hotelier
	result := db.Find(&hotelier)
	if result.Error != nil {
		panic("Error")
	}
	return hotelier
}

// function which work with hotel handler
// get all hotels in system
func GetHotels(db *gorm.DB) []models.Hotel {
	var hotel []models.Hotel
	result := db.Find(&hotel)
	if result.Error != nil {
		panic("Error")
	}
	return hotel
}

// function which work with hotel handler
// get hotel from database by id
func GetHotelByID(db *gorm.DB, id int) models.Hotel {
	var hotel models.Hotel
	result := db.First(&hotel, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("Запись не найдена")
	}
	return hotel
}

// function which work with hotel handler
// This function must use a client
func GetHotelByRating(db *gorm.DB, stars ...int) []models.Hotel {
	var hotels []models.Hotel

	result := db.Where("rating IN ?", stars).Find(hotels)
	if result.Error != nil {
		panic("Error")
	}

	return hotels
}

// function which work with hotel handler
// This function must use a client
func GetHotelByPrice(db *gorm.DB, left, right int) []models.Hotel {
	var hotels []models.Hotel
	result := db.Where("rating > ? AND rating < ?", left, right).Find(hotels)
	if result.Error != nil {
		panic("Error")
	}
	return hotels
}

// function which work with hotel handler
// now any hoteliers can delete any hotel
// function can use only hotelier but delete only his hotels
func DeleteHotelByID(db *gorm.DB, id int) bool {
	result := db.Delete(&models.Hotel{}, id)
	if result.Error != nil {
		panic("Error")
	}
	return true
}

// creating hotels
// function can use only hotelier
func CreateHotel(db *gorm.DB, hotel models.Hotel) {
	db.Create(&hotel)
}

// when registration
func CreateHotelier(db *gorm.DB, hotelier models.Hotelier) {
	db.Create(&hotelier)
}
