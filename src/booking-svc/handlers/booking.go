package handlers

import (
	"github.com/maxturyev/booking-system-project/booking-svc/models"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/maxturyev/booking-system-project/booking-svc/databases"
	"gorm.io/gorm"
)

// Bookings is a http.Handler
type Bookings struct {
	l  *log.Logger
	db *gorm.DB
}

// NewHotels creates a products handler with the given logger
func NewBookings(l *log.Logger, db *gorm.DB) *Bookings {
	return &Bookings{l, db}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface
func (h *Bookings) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the request to get a list of hotels
	if r.Method == http.MethodGet {
		h.getBookings(w)
		return
	}
	// handle the request to add a hotel
	if r.Method == http.MethodPost {
		h.addBooking(w, r)
		return
	}
	if r.Method == http.MethodPut {
		h.UpdateBookingStatus(w, r)
	}
	if r.Method == http.MethodPost {
		h.handleUploadImage(w, r)
		return
	}
}

// getBookings returns the hotels from the date store
func (h *Bookings) getBookings(w http.ResponseWriter) {
	h.l.Println("Handle GET")

	// fetch the hotels from the datastore
	lh := databases.GetBookings(h.db)

	// serialize the list to JSON
	if err := ToJSON(lh, w); err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// methon which can changing any rows in database
func (h *Bookings) UpdateBookingStatus(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle PUT")

	var booking models.Booking

	if err := FromJSON(&booking, r.Body); err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	//checking a correctly seriliazation of format json to our models
	h.l.Println(booking)
	if err := databases.UpdateBooking(h.db, booking); err != nil {
		h.l.Println(err)
	}
}

// addBooking adds a hotel to the date store
func (h *Bookings) addBooking(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle POST")

	var hotel models.Booking

	// deserialize the struct from JSON
	if err := FromJSON(&hotel, r.Body); err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	databases.CreateBooking(h.db, hotel)
}

// POST upload image
func (h *Bookings) handleUploadImage(w http.ResponseWriter, r *http.Request) {
	//Max file size 10MB
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	if err := r.ParseMultipartForm(5 << 10); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.MultipartForm.RemoveAll()
	// Take a file from form
	uf, ufh, err := r.FormFile("media")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer uf.Close()
	//file save in storage in my directory
	flagStoragePath := "C:/Users/rapil/booking-system-project/src/hotel-svc/storage/media"
	path := filepath.Join(flagStoragePath, ufh.Filename)
	if os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//Create a new file in system
	f, err := os.Create(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer f.Close()
	//Copy file in f
	if _, err := io.Copy(f, uf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
