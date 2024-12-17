package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

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
func NewClients(l *log.Logger, db *gorm.DB) *Clients {
	return &Clients{l, db}
}

// ListClients handles GET request to list clients
func (c *Clients) ListClients(ctx *gin.Context) {
	c.l.Println("Handle GET clients")

	// fetch the hotels from the datastore
	lh := databases.GetClients(c.db)

	// serialize the list to JSON
	ctx.JSON(http.StatusOK, lh)
}

// UpdateClient handles PUT request to update a client
func (c *Clients) UpdateClient(ctx *gin.Context) {
	c.l.Println("Handle PUT client")

	var client models.Client

	ctx.JSON(http.StatusOK, &client)

	// checking serialization
	c.l.Println(client)

	if err := databases.UpdateClient(c.db, client); err != nil {
		c.l.Println(err)
	}
}

// AddClient handles POST request to add a client
func (c *Clients) AddClient(ctx *gin.Context) {
	c.l.Println("Handle POST client")

	var client models.Client

	// deserialize the struct from JSON
	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	databases.AddClient(c.db, client)
}
