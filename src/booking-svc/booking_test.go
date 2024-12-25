package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWork(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9091/")
	if err != nil {
		log.Println(err)
		assert.Equal(t, false, true)
		return
	}
	defer resp.Body.Close()
	assert.Equal(t, true, true)
}

func TestBookingGet(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9091/booking/")
	if err != nil {
		log.Println(err)
		assert.Equal(t, false, true)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, 200)
}

func TestBookingHotelGet(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9091/booking/hotel/")
	if err != nil {
		log.Println(err)
		assert.Equal(t, false, true)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, 200)
}

func TestClientGet(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9091/client/")
	if err != nil {
		log.Println(err)
		assert.Equal(t, false, true)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, 200)
}
