package handlers

import (
	"log"
	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"
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
		ctx.JSON(503, gin.H{"error": "Bad news"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"answer": "Good news, everything is fine"})
	}

	// // fetch the hotels from the database
	// lh := postgres.SelectHotels(h.postgres)

	// // serialize the list to JSON
	// ctx.JSON(http.StatusOK, lh)
}

// // GetHotelByID handles GET request to return a hotel by id
// func (h *Hotels) GetHotelByID(ctx *gin.Context) {
// 	h.l.Println("Handle GET")
// 	id, _ := strconv.Atoi(ctx.Param("id"))

// 	// fetch the hotel from the database
// 	hotel := postgres.SelectHotelByID(h.postgres, id)

// 	// serialize the model to JSON
// 	ctx.JSON(http.StatusOK, hotel)
// }

// // PutHotel handles PUT request to update a hotel
// func (h *Hotels) PutHotel(ctx *gin.Context) {
// 	h.l.Println("Handle PUT")

// 	var hotel models.Hotel

// 	// deserialize the struct from JSON
// 	if err := ctx.ShouldBindJSON(&hotel); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := postgres.UpdateHotel(h.postgres, hotel); err != nil {
// 		h.l.Println(err)
// 	}
// }

// // PostHotel handles POST request to create a hotel
// func (h *Hotels) PostHotel(ctx *gin.Context) {
// 	h.l.Println("Handle POST")

// 	var hotel models.Hotel

// 	// deserialize the struct from JSON
// 	if err := ctx.ShouldBindJSON(&hotel); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	}

// 	postgres.CreateHotel(h.postgres, hotel)
// }

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
