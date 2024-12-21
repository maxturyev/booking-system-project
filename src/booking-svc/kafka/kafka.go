package kafka

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/src/booking-svc/models"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
	"os"
	"strconv"
)

const KafkaTopic = "bookings"

// ConnectKafka establishes a connection to kafka
func ConnectKafka() (*kafka.Conn, error) {
	// Setting up a kafka producer
	address := os.Getenv("KAFKA_ADDRESS")
	topic := KafkaTopic
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	return conn, nil
}

// SendMessage sends booking info to notification service
func SendMessage(kc *kafka.Conn, booking models.Booking) error {
	bookingJSON, err := json.Marshal(booking)
	if err != nil {
		log.Println(http.StatusInternalServerError, gin.H{"error": "failed to marshal JSON"})
		return err
	}

	msg := kafka.Message{
		Topic: KafkaTopic,
		Key:   []byte(strconv.Itoa(booking.ClientID)),
		Value: bookingJSON,
	}

	// Send message to kafka
	_, err = kc.WriteMessages(msg)
	if err != nil {
		log.Println("failed to write messages:", err)
		return err
	}

	return nil
}
