package handlers

import (
	"log"
	"net/http"
)

type Hoteliers struct {
	l *log.Logger
}

func NewHoteliers(l *log.Logger) *Hoteliers {
	return &Hoteliers{l}
}

func (h *Hoteliers) ServeHTTP(w http.Response, r *http.Request) {

}
