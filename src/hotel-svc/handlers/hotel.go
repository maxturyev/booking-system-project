package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/hotel-svc/databases"
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

// getHotels returns the hotels from the date store
func (h *Hotels) GetHotels(ctx *gin.Context) {
	h.l.Println("Handle GET")

	// fetch the hotels from the datastore
	lh := databases.GetHotels(h.db)

	ctx.JSON(http.StatusOK, lh)
}

func (h *Hotels) GetHotelByID(ctx *gin.Context) {
	h.l.Println("Handle GET")
	id, _ := strconv.Atoi(ctx.Param("id"))
	lh := databases.GetHotelByID(h.db, id)

	ctx.JSON(http.StatusOK, lh)
}

// methon which can changing any rows in database
func (h *Hotels) UpdateHotel(ctx *gin.Context) {
	h.l.Println("Handle PUT")

	var hotel models.Hotel

	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//checking a correctly serilization of format json to our models
	h.l.Println(hotel)
	err := databases.UpdateHotel(h.db, hotel)
	h.l.Println(err)
}

// addHotel adds a hotel to the date store
func (h *Hotels) AddHotel(ctx *gin.Context) {
	h.l.Println("Handle POST")

	var hotel models.Hotel

	// deserialize the struct from JSON
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	databases.CreateHotel(h.db, hotel)
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
