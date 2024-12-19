package main

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/booking-svc/common"
	pb "github.com/maxturyev/booking-system-project/src/grpc"
	"google.golang.org/grpc"

	"github.com/maxturyev/booking-system-project/booking-svc/db"
	"github.com/maxturyev/booking-system-project/booking-svc/handlers"
	"github.com/segmentio/kafka-go"
)

func validateNumericID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		match, _ := regexp.MatchString(`^\d+$`, id)
		if !match {
			c.JSON(http.StatusBadRequest, gin.H{"error": "non numeric id"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	// Load envs
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// to produce messages
	topic := "my-topic"
	partition := 0

	// Init kafka connection
	kafkaConn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	kafkaConn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	// Generate http server config
	cfg := common.NewConfig()

	// Create logger
	l := log.New(os.Stdout, "booking-svc", log.LstdFlags)

	// Connect to database
	bookingDb := db.ConnectDB()

	// Grpc client server connection
	conn, err := grpc.NewClient(os.Getenv("HOTEL_SERVER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	grpcClient := pb.NewHotelServiceClient(conn)

	router := gin.Default()

	bh := handlers.NewBookings(l, bookingDb, kafkaConn)
	ch := handlers.NewClients(l, bookingDb)

	// Handle requests for booking
	bookingGroup := router.Group("/booking")
	{
		// Handler to get room price of a hotel by ID
		bookingGroup.GET("/hotel", bh.GetHotels(grpcClient))
		bookingGroup.GET("/hotel/:id", validateNumericID(), bh.GetHotelPriceByID(grpcClient))
		bookingGroup.GET("/", bh.GetBookings)
		bookingGroup.POST("/", bh.PostBooking)
		bookingGroup.PUT("/", bh.PutBooking)

	}

	// Handle requests for client
	clientGroup := router.Group("/client")
	{
		clientGroup.GET("/", ch.GetClients)
		clientGroup.POST("/", ch.PostClient)
		clientGroup.PUT("/", ch.UpdateClient)
	}

	// Set up a channel to listen to for interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		cfg.Server.Timeout.Server,
	)
	defer cancel()

	// Define server options
	server := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      router,
		ErrorLog:     l,
		ReadTimeout:  cfg.Server.Timeout.Read * time.Second,
		WriteTimeout: cfg.Server.Timeout.Write * time.Second,
		IdleTimeout:  cfg.Server.Timeout.Idle * time.Second,
	}

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	// Alert the user that the server is starting
	log.Printf("Server is starting on %s\n", server.Addr)

	// Run the server on a new goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				// Normal interrupt operation, ignore
			} else {
				log.Fatalf("Server failed to start due to err: %v", err)
			}
		}
	}()

	// Block on this channel listening for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	log.Printf("Server is shutting down due to %+v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}

}
