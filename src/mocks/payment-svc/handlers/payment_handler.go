package handlers

import (
	"log"
	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
)

// Hotels is a http.Handler
type HandlerPayments struct {
	l  *log.Logger
	db *gorm.DB
}

// NewHotels creates a hotels handler
func NewPayments(l *log.Logger, db *gorm.DB) *HandlerPayments {
	return &HandlerPayments{l, db}
}

// GetHotels handles GET request to list all hotels
func (h *HandlerPayments) ReturnError(ctx *gin.Context) {
	h.l.Println("Now is some error 5xx")

	variant := rand.Intn(3)
	if variant < 2 {
		ctx.JSON(500, gin.H{"answer": "Bad news"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"answer": "Good news, everything is fine"})
	}

	// // fetch the hotels from the database
	// lh := db.SelectHotels(h.db)

	// // serialize the list to JSON
	// ctx.JSON(http.StatusOK, lh)
}

func (h *HandlerPayments) DoPrometeus(ctx *gin.Context) {
	h.l.Println("Prometeus:")

	ctx.JSON(http.StatusOK, promhttp.Handler())

	// // fetch the hotels from the database
	// lh := db.SelectHotels(h.db)

	// // serialize the list to JSON
	// ctx.JSON(http.StatusOK, lh)
}
