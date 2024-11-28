
package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/maxturyev/booking-system-project/api/data"
)

// Hotels is a http.Handler
type Hotels struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewHotels(l *log.Logger) *Hotels {
	return &Hotels{l}
}

// ServeHTTP is the main entry point for the handler and satisfires the http.Handler interface
func (h *Hotels) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the request to get a list of hotels
	if r.Method == http.MethodGet {
		reg := regexp.MustCompile(`/hotels/([0-9]+)`)
		matches := reg.FindStringSubmatch(r.URL.Path)
		if len(matches) != 1 {
			hotelID, _ := strconv.Atoi(matches[1])
			h.GetHotelsByID(w, r, hotelID)
			return
		}
		if r.URL.Path == "/hotels" {
			h.getHotels(w)
			return
		}
	}

	// handle the request to add a hotel
	if r.Method == http.MethodPost {
		h.addHotel(w, r)
		return
	}
}

// getHotels returns the hotels from the date store
func (h *Hotels) getHotels(w http.ResponseWriter) {
	h.l.Println("Handle GET")

	// fetch the hotels from the datastore
	lh := data.GetHotels()

	// serialize the list to JSON
	if err := data.ToJSON(lh, w); err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}

}

func (h *Hotels) GetHotelsByID(w http.ResponseWriter, r *http.Request, ID int) {
	h.l.Println("Handle GET")
	for _, curr := range data.GetHotels() {
		if curr.ID == ID {
			data.ToJSON(*curr, w)
		}
	}
	http.Error(w, "This ID do not exist", http.StatusNotFound)
}

func (h *Hotels) changeNameHotel(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle POST")

}

// addHotel adds a hotel to the date store
func (h *Hotels) addHotel(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle POST")

	hotel := &data.Hotel{}

	// deserialize the struct from JSON
	if err := data.FromJSON(hotel, r.Body); err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	// add a hotel to the data store
	data.AddHotel(hotel)
}
