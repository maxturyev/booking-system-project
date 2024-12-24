package main

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"github.com/maxturyev/booking-system-project/src/booking-svc/kafka"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/src/booking-svc/common"

	pb "github.com/maxturyev/booking-system-project/src/grpc"
	"google.golang.org/grpc"

	"github.com/maxturyev/booking-system-project/src/booking-svc/handlers"
	"github.com/maxturyev/booking-system-project/src/booking-svc/postgres"
)

func main() {
	// Load envs
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Generate http server config
	cfg := common.NewConfig()

	// Create logger
	l := log.New(os.Stdout, "booking-svc", log.LstdFlags)

	// Connect to database
	bookingDb := postgres.ConnectDB()

	// Init kafka connection
	kafkaConn, err := kafka.ConnectKafka()

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
		bookingGroup.GET("/", bh.GetBookings)
		bookingGroup.POST("/", bh.PostBooking)
		bookingGroup.PUT("/", bh.PutBooking)
		bookingGroup.GET("/hotel", bh.GetHotels(grpcClient))
		bookingGroup.GET("/hotel/:id", bh.ValidateNumericID(), bh.GetHotelPriceByID(grpcClient))
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
