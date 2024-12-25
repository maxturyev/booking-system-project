package handlers

import (
	"log"
	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandlerPayments is a http.Handler
type HandlerPayments struct {
	l  *log.Logger
	db *gorm.DB
}

// NewPayments creates a hotels handler
func NewPayments(l *log.Logger, db *gorm.DB) *HandlerPayments {
	return &HandlerPayments{l, db}
}

// ReturnError handles GET request to get 500 or 200 http error
func (h *HandlerPayments) ReturnError(ctx *gin.Context) {
	h.l.Println("Now is some error 5xx")

	variant := rand.Intn(3)
	if variant < 2 {
		ctx.JSON(500, gin.H{"answer": "Bad news"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"answer": "Good news, everything is fine"})
	}
}
