package handlers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/src/hotel-svc/db"
	"github.com/maxturyev/booking-system-project/src/hotel-svc/models"
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
	lh, err := db.SelectHoteliers(h.db)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
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
	res := checkCorrectFieldHotelier(hotelier)
	if res == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not correct field"})
	}
	err := db.CreateHotelier(h.db, hotelier)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func checkCorrectFieldHotelier(hotelier models.Hotelier) bool {
	name, _ := regexp.MatchString(`^[A-Z][a-z]{0,30}$`, hotelier.FirstName)
	secondname, _ := regexp.MatchString(`^[A-Z][a-z]{0,30}$`, hotelier.LastName)
	email, _ := regexp.MatchString(`/^((([0-9A-Za-z]{1}[-0-9A-z\.]{1,}[0-9A-Za-z]{1})|([0-9А-Яа-я]{1}[-0-9А-я\.]{1,}[0-9А-Яа-я]{1}))@([-A-Za-z]{1,}\.){1,2}[-A-Za-z]{2,})$`,
		hotelier.Email)
	phone, _ := regexp.MatchString(`^\+[7][0-9]{10}$`, hotelier.Phone)
	login, _ := regexp.MatchString(`[a-zA-Z0-9]{,30}`, hotelier.Phone)
	return name && secondname && email && phone && login
}
