package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

const (
	KafkaTopic = "bookings"
)

func main() {
	// to consume messages
	address := os.Getenv("KAFKA_SERVER_ADDR")
	topic := KafkaTopic
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	for {
		batch := conn.ReadBatch(1e3, 1e6) // fetch 10KB min, 1MB max

		b := make([]byte, 10e3) // 10KB max per message
		for {
			n, err := batch.Read(b)
			if err != nil {
				break
			}
			fmt.Println(string(b[:n]))
		}
		if err := batch.Close(); err != nil {
			log.Fatal("failed to close batch:", err)
		}
	}
}
