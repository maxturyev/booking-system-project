package main

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWork(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9095/booking/")
	if err != nil {
		log.Println(err)
		assert.Equal(t, false, true)
		return
	}
	defer resp.Body.Close()
	assert.Equal(t, true, true)
}

func TestBookingGet(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9095/booking/")
	if err != nil {
		log.Println(err)
		assert.Equal(t, false, true)
		return
	}
	defer resp.Body.Close()
	answer := map[string]string{}

	log.Println(answer)
	for true {
		bs := make([]byte, 1024)
		n, err := resp.Body.Read(bs)
		log.Println(string(bs[:n]))

		json.Unmarshal([]byte(string(bs[:n])), &answer)

		log.Println(answer["answer"])

		if n == 0 || err != nil {
			break
		}
	}

	flag := false
	if answer["answer"] == "Handle GET bookings" {
		flag = true
	}

	assert.Equal(t, flag, true)
}
