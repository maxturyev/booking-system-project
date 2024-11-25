package handlers

import (
	"log"
	"net/http"

	"github.com/maxturyev/booking-system-project/api/data"
)

// Hotels is a http.Handler
type Client struct {
	l *log.Logger
}

// NewHotels creates a products handler with the given logger
func NewClient(l *log.Logger) *Client {
	return &Client{l}
}

// ServeHTTP is the main entry point for the handler and satisfires the http.Handler interface
func (c *Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the request to get a list of hotels
	if r.Method == http.MethodGet {
		switch r.URL.Path {
		case "/client/hotels":
			c.getHotels(w, r)
		case "/client/bookings":
			c.getBookings(w, r)
		}
	}

	// handle the request to add a hotel
	if r.Method == http.MethodPost {
		c.createBooking(w, r)
		return
	}
}

// getHotels returns the hotels from the date store
func (c *Client) getHotels(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle GET hotels")

	// fetch the hotels from the datastore
	lh := data.GetHotels()

	// serialize the list to JSON
	if err := data.ToJSON(lh, w); err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// createBooking creates a new booking
func (c *Client) createBooking(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle POST")

	booking := &data.Booking{}

	// deserialize the struct from JSON
	if err := data.FromJSON(booking, r.Body); err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	for _, hotel := range data.GetHotels() {
		if hotel.ID == booking.HotelID && hotel.RoomsAvailable > 0 {
			data.CreateBooking(booking)
			hotel.RoomsAvailable -= 1
		}
	}
}

// getHotels returns the hotels from the date store
func (c *Client) getBookings(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle GET bookings")

	// fetch the hotels from the datastore
	lb := data.GetBookings()

	// serialize the list to JSON
	if err := data.ToJSON(lb, w); err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
