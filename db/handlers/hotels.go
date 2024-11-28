package handlers

import (
	"fmt"
	"time"

	"github.com/maxturyev/booking-system-project/consts"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var hotel_db *gorm.DB

// Hotel defines the structure for an API hotel
type Hotel struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Hotelier       int    `json:"hotelier_id"`
	Rating         int    `json:"rating"`
	Country        string `json:"country"`
	Address        string `json:"address"`
	RoomsAvailable int    `json:"rooms_available"`
	CreatedOn      string `json:"-"`
	UpdatedOn      string `json:"-"`
	DeletedOn      string `json:"-"`
}

// Hotels is a collection of Product
type Hotels []*Hotel

func NewHotelConnection() error {
	// Initialize connection to Hotel database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", consts.Host, consts.Port, consts.User, consts.Password, consts.DBHotel)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	hotel_db = db

	return nil
}

// GetHotels returns a list of hotels
func GetHotels() Hotels {
	return hotelList
}

func InsertSingle(tableName string, hotel *Hotel) {
	hotel_db.Table(tableName).Create(hotel)
}

// Example data store
var hotelList = Hotels{
	{
		ID:             1,
		Name:           "Hilton Moscow Leningradskaya",
		Hotelier:       1,
		Rating:         5,
		Country:        "Russia",
		Address:        "Kalanchevskaya street 21/40",
		RoomsAvailable: 123,
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
	{
		ID:             2,
		Name:           "Cavalieri Hotel",
		Hotelier:       1,
		Rating:         4,
		Country:        "Greece",
		Address:        "Kapodistriou street 4",
		RoomsAvailable: 0,
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
}
