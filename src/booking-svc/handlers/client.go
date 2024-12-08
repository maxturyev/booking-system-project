package handlers

import (
	"log"
	"net/http"

	"github.com/maxturyev/booking-system-project/booking-svc/models"

	"github.com/maxturyev/booking-system-project/booking-svc/databases"
	"gorm.io/gorm"
)

// Clients is a http.Handler
type Clients struct {
	l  *log.Logger
	db *gorm.DB
}

// NewClients creates a clients handler with the given logger
func NewClients(l *log.Logger, db *gorm.DB) *Bookings {
	return &Bookings{l, db}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface
func (c *Clients) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the request to get a list of hotels
	if r.Method == http.MethodGet {
		c.getClients(w)
		return
	}

	// handle the request to add a hotel
	if r.Method == http.MethodPost {
		c.addClient(w, r)
		return
	}
	// handle the request to update client info
	if r.Method == http.MethodPut {
		c.updateClient(w, r)
	}
}

// getBookings returns the hotels from the date store
func (c *Clients) getClients(w http.ResponseWriter) {
	c.l.Println("Handle GET clients")

	// fetch the hotels from the datastore
	lh := databases.GetClients(c.db)

	// serialize the list to JSON
	if err := ToJSON(lh, w); err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// UpdateClient changes client info
func (c *Clients) updateClient(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle PUT client")

	var booking models.Booking

	if err := FromJSON(&booking, r.Body); err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	// checking serialization
	c.l.Println(booking)

	if err := databases.UpdateBooking(c.db, booking); err != nil {
		c.l.Println(err)
	}
}

// addClient adds a client to the database
func (c *Clients) addClient(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle POST client")

	var client models.Client

	// deserialize the struct from JSON
	if err := FromJSON(&client, r.Body); err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	databases.AddClient(c.db, client)
}
