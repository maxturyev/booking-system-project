package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/hotel-svc/db"
	"github.com/maxturyev/booking-system-project/hotel-svc/models"
	"gorm.io/gorm"
)

// Hoteliers is a http.Handler
type Hoteliers struct {
	l  *log.Logger
	db *gorm.DB
}

// NewHoteliers creates a hoteliers handler
func NewHoteliers(l *log.Logger, db *gorm.DB) *Hoteliers {
	return &Hoteliers{l, db}
}

// GetHoteliers handles GET request to list all hoteliers
func (h *Hoteliers) GetHoteliers(ctx *gin.Context) {
	h.l.Println("Handle GET Hoteliers")

	// fetch the hoteliers from the database
	lh := db.SelectHoteliers(h.db)

	// serialize the list to JSON
	ctx.JSON(http.StatusOK, lh)
}

// PostHotelier handles POST request to create a hotelier
func (h *Hoteliers) PostHotelier(ctx *gin.Context) {
	h.l.Println("Handle POST Hotelier")

	var hotelier models.Hotelier

	// deserialize the struct from JSON
	if err := ctx.ShouldBind(&hotelier); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	db.CreateHotelier(h.db, hotelier)
}
