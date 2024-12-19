package handlers

import (
	"context"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxturyev/booking-system-project/booking-svc/db"
	"github.com/maxturyev/booking-system-project/booking-svc/models"
	pb "github.com/maxturyev/booking-system-project/src/grpc"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

// Bookings is a http.Handler
type Bookings struct {
	l  *log.Logger
	db *gorm.DB
	kc *kafka.Conn
}

// NewBookings creates a bookings handler
func NewBookings(l *log.Logger, db *gorm.DB, broker *kafka.Conn) *Bookings {
	return &Bookings{l, db, broker}
}

// GetBookings handles GET request to list all bookings
func (c *Bookings) GetBookings(ctx *gin.Context) {
	c.l.Println("Handle GET bookings")

	// fetch the hotels from the database
	lh := db.SelectBookings(c.db)

	// serialize the list to JSON
	ctx.JSON(http.StatusOK, lh)
}

// PutBooking handles PUT request to update a booking
func (c *Bookings) PutBooking(ctx *gin.Context) {
	c.l.Println("Handle PUT")

	var booking models.Booking

	// deserialize http request body
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := db.UpdateBooking(c.db, booking); err != nil {
		c.l.Println(err)
	}
}

// PostBooking handles a POST request to create a booking
func (c *Bookings) PostBooking(ctx *gin.Context) {
	c.l.Println("Handle POST")

	var booking models.Booking

	// deserialize the struct from JSON
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// BookingEvent - структура для нашего сообщения
	type BookingEvent struct {
		ID    uint32 `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	requestID := uuid.New().ID()

	bookingEvent := BookingEvent{
		ID: requestID,

	}
	_, err := c.kc.WriteMessages(
		kafka.Message{Key: uuid., Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := kafkaConn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	db.CreateBooking(c.db, booking)
}

func (c *Bookings) GetHotelPriceByID(grpcClient pb.HotelServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Context with the amount of time to process the grpc request
		ctxgrpc, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		log.Println("Grpc connection established")
		id, _ := strconv.Atoi(ctx.Param("id"))
		log.Println(id)
		response, err := grpcClient.GetHotelPriceByID(ctxgrpc, &pb.GetHotelPriceByIDRequest{Id: int32(id)})
		if err != nil {
			log.Fatal("error in grpc get price")
		}

		ctx.JSON(200, gin.H{
			"hotel price": response.RoomPrice,
		})
	}
}

func (c *Bookings) GetHotels(grpcClient pb.HotelServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Context with the amount of time to process the grpc request
		ctxgrpc, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		log.Println("Handle GRPC request to get hotel room price")

		stream, err := grpcClient.GetHotels(ctxgrpc, &pb.GetHotelsRequest{})
		if err != nil {
			log.Fatalf("Error setting up a stream %v", err)
		}

		var hotelList []struct {
			HotelID        uint   `json:"hotel_id"`
			Name           string `json:"name"`
			Rating         int    `json:"rating"`
			Country        string `json:"country"`
			Description    string `json:"description"`
			RoomsAvailable int    `json:"rooms_available"`
			Price          int    `json:"price"`
			Address        string `json:"address"`
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Print("Grpc connection ended")
				break
			}
			if err != nil {
				log.Print("error")
				ctx.JSON(500, gin.H{"error": "Error from getting information"})
				return
			}
			hotelList = append(hotelList, struct {
				HotelID        uint   `json:"hotel_id"`
				Name           string `json:"name"`
				Rating         int    `json:"rating"`
				Country        string `json:"country"`
				Description    string `json:"description"`
				RoomsAvailable int    `json:"rooms_available"`
				Price          int    `json:"price"`
				Address        string `json:"address"`
			}{
				HotelID:        uint(res.Hotel.HotelID),
				Name:           res.Hotel.Name,
				Rating:         int(res.Hotel.Rating),
				Country:        res.Hotel.Country,
				Description:    res.Hotel.Description,
				RoomsAvailable: int(res.Hotel.RoomAvailable),
				Price:          int(res.Hotel.RoomPrice),
				Address:        res.Hotel.Address,
			})
		}

		ctx.JSON(200, gin.H{
			"hotels": hotelList,
		})
	}
}
