package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/hotel-svc/db"
	"github.com/maxturyev/booking-system-project/hotel-svc/models"
	"gorm.io/gorm"
)

type Hotelier struct {
	l  *log.Logger
	db *gorm.DB
}

func NewHotelier(l *log.Logger, db *gorm.DB) *Hotelier {
	return &Hotelier{l, db}
}

func (h *Hotelier) GetHoteliers(ctx *gin.Context) {
	h.l.Println("Handle GET")
	lh := db.GetHoteliers(h.db)
	ctx.JSON(http.StatusOK, lh)
}

func (h *Hotelier) AddHotel(ctx *gin.Context) {
	h.l.Println("Handle POST")
	var hotelier models.Hotelier

	if err := ctx.ShouldBind(&hotelier); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	db.CreateHotelier(h.db, hotelier)
}
