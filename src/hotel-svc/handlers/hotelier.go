// This file must have a authorize func which checking a permission
package handlers

import (
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Hoteliers struct {
	l  log.Logger
	db *gorm.DB
}

func NewHotelier(l *log.Logger, db *gorm.DB) *Hoteliers {
	return &Hoteliers{}
}

func (h *Hoteliers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//handle the request to get a list of hoteliers

}
