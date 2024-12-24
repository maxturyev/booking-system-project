package main

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentGet(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9095/payment/")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	answer := map[string]string{}
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
	if answer["answer"] == "Good news, everything is fine" || answer["answer"] == "Bad news" {
		flag = true
	}

	assert.Equal(t, flag, true)
}
