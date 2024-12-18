package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/hotel-svc/db"
	"github.com/maxturyev/booking-system-project/hotel-svc/models"
	"gorm.io/gorm"
)

// Hotels is a http.Handler
type Hotels struct {
	l  *log.Logger
	db *gorm.DB
}

// NewHotels creates a products handler with the given logger
func NewHotels(l *log.Logger, db *gorm.DB) *Hotels {
	return &Hotels{l, db}
}

// GetHotels returns the hotels from the database
func (h *Hotels) GetHotels(ctx *gin.Context) {
	h.l.Println("Handle GET")

	// fetch the hotels from the datastore
	lh := db.GetHotels(h.db)

	ctx.JSON(http.StatusOK, lh)
}

// GetHotelByID returns the hotels from the database by ID
func (h *Hotels) GetHotelByID(ctx *gin.Context) {
	h.l.Println("Handle GET")
	id, _ := strconv.Atoi(ctx.Param("id"))
	lh := db.GetHotelByID(h.db, id)

	ctx.JSON(http.StatusOK, lh)
}

// UpdateHotel updates hotel info
func (h *Hotels) UpdateHotel(ctx *gin.Context) {
	h.l.Println("Handle PUT")

	var hotel models.Hotel

	// deserialize the struct from JSON
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := db.UpdateHotel(h.db, hotel)
	h.l.Println(err)
}

// AddHotel adds a hotel to the database
func (h *Hotels) AddHotel(ctx *gin.Context) {
	h.l.Println("Handle POST")

	var hotel models.Hotel

	// deserialize the struct from JSON
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	db.CreateHotel(h.db, hotel)
}

// // POST upload image
// func (h *Hotels) HandleUploadImage(ctx *gin.Context) {
// 	//Max file size 10MB
// 	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
// 	if err := r.ParseMultipartForm(5 << 10); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	defer r.MultipartForm.RemoveAll()
// 	// Take a file from form
// 	uf, ufh, err := r.FormFile("media")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	defer uf.Close()
// 	//file save in storage in my directory
// 	flagStoragePath := "C:/Users/rapil/booking-system-project/src/hotel-svc/storage/media"
// 	path := filepath.Join(flagStoragePath, ufh.Filename)
// 	if os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	//Create a new file in system
// 	f, err := os.Create(path)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer f.Close()
// 	//Copy file in f
// 	if _, err := io.Copy(f, uf); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
