package funcs

import (
	"github.com/maxturyev/booking-system-project/db/handlers"
)

// AddHotel adds a new hotel to the data store
func AddHotel(h *handlers.Hotel) {
	handlers.InsertSingle("hotel", h)
}
