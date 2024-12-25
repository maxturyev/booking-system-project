package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWork(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9090/")
	if err != nil {
		log.Println(err)
		assert.Equal(t, false, true)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, true, true)
}

func TestHotelGet(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9090/hotel/")
	if err != nil {
		log.Println(err)
		assert.Equal(t, false, true)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, 200)
}

func TestHotelierGet(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9090/hotelier/")
	if err != nil {
		log.Println(err)
		assert.Equal(t, false, true)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, 200)
}
