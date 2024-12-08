package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/maxturyev/booking-system-project/databases"
	"github.com/maxturyev/booking-system-project/models"
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

// ServeHTTP is the main entry point for the handler and satisfires the http.Handler interface
func (h *Hotels) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the request to get a list of hotels
	if r.Method == http.MethodGet {
		h.getHotels(w)
		return
	}
	// handle the request to add a hotel
	if r.Method == http.MethodPost {
		h.addHotel(w, r)
		return
	}
	if r.Method == http.MethodPut {
		h.UpdateHotel(w, r)
	}
	if r.Method == http.MethodPost {
		h.handleUploadImage(w, r)
		return
	}
}

// getHotels returns the hotels from the date store
func (h *Hotels) getHotels(w http.ResponseWriter) {
	h.l.Println("Handle GET")

	// fetch the hotels from the datastore
	lh := databases.GetHotels(h.db)

	// serialize the list to JSON
	if err := ToJSON(lh, w); err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// methon which can changing any rows in database
func (h *Hotels) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle PUT")

	var hotel models.Hotel

	if err := FromJSON(&hotel, r.Body); err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}
	//checking a correctly serilization of format json to our models
	h.l.Println(hotel)
	err := databases.UpdateHotel(h.db, hotel)
	h.l.Println(err)
}

// addHotel adds a hotel to the date store
func (h *Hotels) addHotel(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle POST")

	var hotel models.Hotel

	// deserialize the struct from JSON
	if err := FromJSON(&hotel, r.Body); err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	databases.CreateHotel(h.db, hotel)
}

// POST upload image
func (h *Hotels) handleUploadImage(w http.ResponseWriter, r *http.Request) {
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
