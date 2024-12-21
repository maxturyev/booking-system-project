package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/maxturyev/booking-system-project/src/booking-svc/models"

	"github.com/maxturyev/booking-system-project/src/booking-svc/postgres"
	"gorm.io/gorm"
)

// Clients is a http.Handler
type Clients struct {
	l  *log.Logger
	db *gorm.DB
}

// NewClients creates a clients handler
func NewClients(l *log.Logger, db *gorm.DB) *Clients {
	return &Clients{l, db}
}

// GetClients handles GET request to list all clients
func (c *Clients) GetClients(ctx *gin.Context) {
	c.l.Println("Handle GET clients")

	// fetch the hotels from the database
	lh := postgres.SelectClients(c.db)

	// serialize the list to JSON
	ctx.JSON(http.StatusOK, lh)
}

// UpdateClient handles PUT request to update a client
func (c *Clients) UpdateClient(ctx *gin.Context) {
	c.l.Println("Handle PUT client")

	var client models.Client

	// deserialize the struct from JSON
	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := postgres.UpdateClient(c.db, client); err != nil {
		c.l.Println(err)
	}
}

// PostClient handles POST request to create a client
func (c *Clients) PostClient(ctx *gin.Context) {
	c.l.Println("Handle POST client")

	var client models.Client

	// deserialize the struct from JSON
	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	postgres.CreateClient(c.db, client)
}
