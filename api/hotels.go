package api

import (
	"log"
	"net/http"

	"github.com/maxturyev/booking-system-project/db"
	"github.com/maxturyev/booking-system-project/db/funcs"
	"github.com/maxturyev/booking-system-project/db/handlers"
)

// Hotels is a http.Handler
type Hotels struct {
	l *log.Logger
}

// NewHotels creates a products handler with the given logger
func NewHotels(l *log.Logger) *Hotels {
	return &Hotels{l}
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
	lh := handlers.GetHotels()

	// serialize the list to JSON
	if err := db.ToJSON(lh, w); err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// addHotel adds a hotel to the date store
func (h *Hotels) addHotel(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle POST")

	hotel := &handlers.Hotel{}

	// deserialize the struct from JSON
	if err := db.FromJSON(hotel, r.Body); err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	// add a hotel to the data store
	funcs.AddHotel(hotel)
}
