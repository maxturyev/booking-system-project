package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/booking-svc/databases"
	"github.com/maxturyev/booking-system-project/booking-svc/models"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Bookings is a http.Handler
type Bookings struct {
	l  *log.Logger
	db *gorm.DB
}

// NewHotels creates a products handler with the given logger
func NewBookings(l *log.Logger, db *gorm.DB) *Bookings {
	return &Bookings{l, db}
}

// ListBookings handles GET request to list bookings
func (c *Bookings) ListBookings(ctx *gin.Context) {
	c.l.Println("Handle GET bookings")

	// fetch the hotels from the database
	lh := databases.GetBookings(c.db)

	// serialize the list to JSON
	ctx.JSON(http.StatusOK, lh)
}

// UpdateBooking handles PUT request to update bookings
func (c *Bookings) UpdateBooking(ctx *gin.Context) {
	c.l.Println("Handle PUT")

	var booking models.Booking

	// deserialize http request body
	ctx.JSON(http.StatusOK, booking)

	// check
	c.l.Println(booking)

	if err := databases.UpdateBooking(c.db, booking); err != nil {
		c.l.Println(err)
	}
}

// CreateBooking handles a POST request to create a booking
func (c *Bookings) CreateBooking(ctx *gin.Context) {
	c.l.Println("Handle POST")

	var hotel models.Booking

	// deserialize the struct from JSON
	ctx.JSON(http.StatusOK, hotel)

	databases.CreateBooking(c.db, hotel)
}
