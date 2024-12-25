package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/maxturyev/booking-system-project/src/grpc"
	"github.com/maxturyev/booking-system-project/src/hotel-svc/common"
	grpcserver "github.com/maxturyev/booking-system-project/src/hotel-svc/grpc-server"
	"github.com/maxturyev/booking-system-project/src/hotel-svc/handlers"
	"github.com/maxturyev/booking-system-project/src/hotel-svc/postgres"
	"google.golang.org/grpc"
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
	// Generate httpServer config
	cfg := common.NewConfig()

	// Create logger
	l := log.New(os.Stdout, "hotel-api", log.LstdFlags)

	// Connect to database
	hotelDb := postgres.ConnectDB()

	// Create grpc httpServer
	grpcServer := grpc.NewServer()
	port := os.Getenv("HOTEL_GRPC_PORT")

	// Run the http httpServer on a new goroutine
	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to create listener: %s", err)
		}

		pb.RegisterHotelServiceServer(grpcServer, &grpcserver.HotelServer{DB: hotelDb})

		log.Printf("Grpc httpServer started  on port %s", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	// Create router and define routes and return that router
	router := gin.Default()

	// Create handlers
	hh := handlers.NewHotels(l, hotelDb)
	hth := handlers.NewHoteliers(l, hotelDb)

	// Handle requests for hotel
	hotelGroup := router.Group("/hotel")
	{
		hotelGroup.GET("/", hh.GetHotels)
		hotelGroup.GET("/:id", validateNumericID(), hh.GetHotelByID)
		hotelGroup.POST("/", hh.PostHotel)
		hotelGroup.POST("/media", hh.HandleUploadImage)
	}

	// Handle requests for hotelier
	hotelierGroup := router.Group("/hotelier")
	{
		hotelierGroup.GET("/", hth.GetHoteliers)
		hotelierGroup.POST("/", hth.PostHotelier)
	}

	// Set up a channel to listen to for interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Set up a context to allow for graceful httpServer shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		cfg.Server.Timeout.Server,
	)
	defer cancel()

	// Define http httpServer options
	httpServer := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.Timeout.Read * time.Second,
		WriteTimeout: cfg.Server.Timeout.Write * time.Second,
		IdleTimeout:  cfg.Server.Timeout.Idle * time.Second,
	}

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	// Alert the user that the httpServer is starting
	log.Printf("Server is starting on %s\n", httpServer.Addr)

	// Run the http httpServer on a new goroutine
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				// Normal interrupt operation, ignore
			} else {
				log.Fatalf("Server failed to start due to err: %v", err)
			}
		}
	}()

	// Block on this channel listening for those previously defined syscalls assign
	// to variable so we can let the user know why the httpServer is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the httpServer
	// while alerting the user
	log.Printf("Server is shutting down due to %+v\n", interrupt)
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}

}
