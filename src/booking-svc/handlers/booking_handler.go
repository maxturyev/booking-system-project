package handlers

import (
	"context"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	kafkaGo "github.com/segmentio/kafka-go"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/src/booking-svc/kafka"
	"github.com/maxturyev/booking-system-project/src/booking-svc/models"
	"github.com/maxturyev/booking-system-project/src/booking-svc/postgres"
	pb "github.com/maxturyev/booking-system-project/src/grpc"
	"gorm.io/gorm"
)

// Bookings is a http.Handler
type Bookings struct {
	l  *log.Logger
	db *gorm.DB
	kc *kafkaGo.Conn
}

// NewBookings creates a bookings handler
func NewBookings(l *log.Logger, db *gorm.DB, kc *kafkaGo.Conn) *Bookings {
	return &Bookings{l, db, kc}
}

// GetBookings handles GET request to list all bookings
func (b *Bookings) GetBookings(ctx *gin.Context) {
	b.l.Println("Handle GET bookings")

	// fetch the hotels from the database
	lh := postgres.SelectBookings(b.db)

	// serialize the list to JSON
	ctx.JSON(http.StatusOK, lh)
}

// PutBooking handles PUT request to update a booking
func (b *Bookings) PutBooking(ctx *gin.Context) {
	b.l.Println("Handle PUT")

	var booking models.Booking

	// deserialize http request body
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Add booking to database
	if err := postgres.UpdateBooking(b.db, booking); err != nil {
		b.l.Println(err)
	}

	// Send kafka message
	// Send kafka message
	if err := kafka.SendMessage(b.kc, booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func getHotelPrice(ctx *gin.Context, grpcClient pb.HotelServiceClient, hotelID int) (float32, error) {
	ctxGrpc, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := grpcClient.GetHotelPriceByID(ctxGrpc, &pb.GetHotelPriceByIDRequest{Id: int32(hotelID)})
	if err != nil {
		return 0, err
	}

	return response.RoomPrice, nil
}

// PostBooking handles a POST request to create a booking
func (b *Bookings) PostBooking(ctx *gin.Context, grpcClient pb.HotelServiceClient) {
	b.l.Println("Handle POST booking")

	var booking models.Booking

	// deserialize the struct from JSON
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Check if "price" and "hotelId" is provided in the input
	if ctx.Request.Body != nil {
		var rawInput map[string]interface{}
		if err := ctx.BindJSON(&rawInput); err == nil {
			if _, existsPrice := rawInput["price"]; existsPrice {
				if _, exitstHotelID := rawInput["hotel_id"]; exitstHotelID {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": "field 'price' is not allowed"})
					return
				}
			}
		}
	}
	id, _ := strconv.Atoi(ctx.Param("hotel_id"))
	price, _ := getHotelPrice(ctx, grpcClient, id)
	booking.HotelID = id
	booking.Price = price

	// Add booking to database
	if err := postgres.CreateBooking(b.db, booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Send kafka message
	if err := kafka.SendMessage(b.kc, booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// ValidateNumericID makes sure that the id parameter is numeric
func (b *Bookings) ValidateNumericID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		match, _ := regexp.MatchString(`^\d+$`, id)
		if !match {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "non numeric id"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// GetHotelPriceByID fetches price of a hotel room from a grpc server (Hotel)
func (b *Bookings) GetHotelPriceByID(grpcClient pb.HotelServiceClient) gin.HandlerFunc {
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

// GetHotels a list of hotels from a grpc server (Hotel)
func (b *Bookings) GetHotels(grpcClient pb.HotelServiceClient) gin.HandlerFunc {
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
