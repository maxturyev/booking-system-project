package handlers

import (
	"log"
	"net/http"

	create "github.com/maxturyev/booking-system-project/src/hotel-svc/databases/hotels_data/CREATE"
	get "github.com/maxturyev/booking-system-project/src/hotel-svc/databases/hotels_data/GET"
	"github.com/maxturyev/booking-system-project/src/hotel-svc/models"
	"gorm.io/gorm"
)

// Hotels is a http.Handler
type Hotels struct {
	l  *log.Logger
	db *gorm.DB
}

// NewHotels creates a products handler with the given logger
func NewHotels(l *log.Logger, db *gorm.DB) *Hotels {
	return &Hotels{l, db}
}

// ServeHTTP is the main entry point for the handler and satisfires the http.Handler interface
func (h *Hotels) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the request to get a list of hotels
	if r.Method == http.MethodGet {
		h.getHotels(w)
		return
	}

	// handle the request to add a hotel
	if r.Method == http.MethodPost {
		h.addHotel(w, r)
		return
	}
}

// getHotels returns the hotels from the date store
func (h *Hotels) getHotels(w http.ResponseWriter) {
	h.l.Println("Handle GET")

	// fetch the hotels from the datastore
	lh := get.GetHotels(h.db)

	// serialize the list to JSON
	if err := ToJSON(lh, w); err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// addHotel adds a hotel to the date store
func (h *Hotels) addHotel(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle POST")

	var hotel models.Hotel

	// deserialize the struct from JSON
	if err := FromJSON(&hotel, r.Body); err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	// add a hotel to the data store
	create.CreateHotel(h.db, hotel)
}
