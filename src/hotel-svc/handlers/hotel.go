package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/src/hotel-svc/models"
	"github.com/maxturyev/booking-system-project/src/hotel-svc/postgres"
	"gorm.io/gorm"
)

// Hotels is a http.Handler
type Hotels struct {
	l  *log.Logger
	db *gorm.DB
}

// NewHotels creates a hotels handler
func NewHotels(l *log.Logger, db *gorm.DB) *Hotels {
	return &Hotels{l, db}
}

// GetHotels handles GET request to list all hotels
func (h *Hotels) GetHotels(ctx *gin.Context) {
	h.l.Println("Handle GET")

	// fetch the hotels from the database
	lh, err := postgres.SelectHotels(h.db)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// serialize the list to JSON
	ctx.JSON(http.StatusOK, lh)
}

func ValidateNumericID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		match, _ := regexp.MatchString(`^\d+$`, id)
		if !match {
			c.JSON(http.StatusBadRequest, gin.H{"error": "non numeric id"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// GetHotelByID handles GET request to return a hotel by id
func (h *Hotels) GetHotelByID(ctx *gin.Context) {
	h.l.Println("Handle GET")
	id, _ := strconv.Atoi(ctx.Param("id"))

	// fetch the hotel from the database
	hotel, err := postgres.SelectHotelByID(h.db, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// serialize the model to JSON
	ctx.JSON(http.StatusOK, hotel)
}

// PutHotel handles PUT request to update a hotel
func (h *Hotels) PutHotel(ctx *gin.Context) {
	h.l.Println("Handle PUT")

	var hotel models.Hotel

	// deserialize the struct from JSON
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := checkCorrectFieldHotel(hotel)
	if res == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not correct field"})
	}
	if err := postgres.UpdateHotel(h.db, hotel); err != nil {
		h.l.Println(err)
	}
}

// PostHotel handles POST request to create a hotel
func (h *Hotels) PostHotel(ctx *gin.Context) {
	h.l.Println("Handle POST hotel")

	var hotel models.Hotel

	// deserialize the struct from JSON
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	res := checkCorrectFieldHotel(hotel)
	if res == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not correct field"})
	}
	err := postgres.CreateHotel(h.db, hotel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// // POST upload image
func (h *Hotels) HandleUploadImage(ctx *gin.Context) {
	//Max file size 10MB
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 10<<20)
	if err := ctx.Request.ParseMultipartForm(5 << 10); err != nil {
		ctx.String(http.StatusBadRequest, "error")
		return
	}
	defer ctx.Request.MultipartForm.RemoveAll()
	// Take a file from form
	uf, ufh, err := ctx.Request.FormFile("media")
	if err != nil {
		ctx.String(http.StatusBadRequest, "error")
		return
	}
	defer uf.Close()
	//file save in storage in my directory
	flagStoragePath := "C:/Users/rapil/booking-system-project/src/hotel-svc/storage/media"
	path := filepath.Join(flagStoragePath, ufh.Filename)
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		ctx.String(http.StatusInternalServerError, "error")
		return
	}
	//Create a new file in system
	f, err := os.Create(path)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "error")
		return
	}
	defer f.Close()
	//Copy file in f
	if _, err := io.Copy(f, uf); err != nil {
		ctx.String(http.StatusInternalServerError, "error")
	}
	ctx.String(http.StatusOK, "ok")
}

func checkCorrectFieldHotel(hotel models.Hotel) bool {
	name, _ := regexp.MatchString(`^[A-Z][a-z]{0,30}$`, hotel.Name)
	address, _ := regexp.MatchString(`^\w{0,30}$`, hotel.Address)
	country, _ := regexp.MatchString(`^[A-Z][a-z]{0,30}$`, hotel.Country)
	return name && address && country
}
