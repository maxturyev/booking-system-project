package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/maxturyev/booking-system-project/src/booking-svc/common"
	"github.com/maxturyev/booking-system-project/src/booking-svc/handlers"
	"github.com/maxturyev/booking-system-project/src/booking-svc/kafka"
	"github.com/maxturyev/booking-system-project/src/booking-svc/postgres"
	pb "github.com/maxturyev/booking-system-project/src/grpc"
)

var (
	requestsTotalBooking = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total_booking",
			Help: "Total number of HTTP requests on booking.",
		},
		[]string{"method"},
	)
	requestDurationBooking = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds_booking",
			Help: "Duration of HTTP requests on booking.",
		},
		[]string{"method"},
	)
	requestsTotalClient = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total_client",
			Help: "Total number of HTTP requests on client.",
		},
		[]string{"method"},
	)
	requestDurationClient = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds_client",
			Help: "Duration of HTTP requests on client.",
		},
		[]string{"method"},
	)
)

func handlerBookingPrometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		start := time.Now()
		elapsed := time.Since(start).Seconds()
		requestsTotalBooking.WithLabelValues(method).Inc()
		requestDurationBooking.WithLabelValues(method).Observe(elapsed)
	}
}

func handlerClientPrometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		start := time.Now()
		elapsed := time.Since(start).Seconds()
		requestsTotalClient.WithLabelValues(method).Inc()
		requestDurationClient.WithLabelValues(method).Observe(elapsed)
	}
}

func prometheusView() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Обработка Прометея
	prometheus.MustRegister(requestsTotalBooking, requestDurationBooking, requestsTotalClient, requestDurationClient)

	// Generate http server config
	cfg := common.NewConfig()

	// Create logger
	l := log.New(os.Stdout, "booking-svc\t", log.LstdFlags)

	// Connect to database
	bookingDb := postgres.ConnectDB()

	// Init kafka connection
	kafkaConn, err := kafka.ConnectKafka()
	if err != nil {
		l.Println("first:", err)
	}

	// Grpc client server connection
	conn, err := grpc.NewClient(os.Getenv("HOTEL_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.Println("second:", err)
	}

	defer conn.Close()

	grpcClient := pb.NewHotelServiceClient(conn)

	router := gin.Default()

	router.GET("/metrics", prometheusView())

	bh := handlers.NewBookings(l, bookingDb, kafkaConn)
	ch := handlers.NewClients(l, bookingDb)

	// Handle requests for booking
	bookingGroup := router.Group("/booking")
	{
		bookingGroup.GET("/", handlerBookingPrometheus(), bh.GetBookings)
		bookingGroup.POST("/", handlerBookingPrometheus(), bh.PostBooking)
		bookingGroup.PUT("/", handlerBookingPrometheus(), bh.PutBooking)
		bookingGroup.GET("/hotel", handlerBookingPrometheus(), bh.GetHotels(grpcClient))
		bookingGroup.GET("/hotel/:id", handlerBookingPrometheus(), bh.ValidateNumericID(), bh.GetHotelPriceByID(grpcClient))
	}

	// Handle requests for client
	clientGroup := router.Group("/client")
	{
		clientGroup.GET("/", handlerClientPrometheus(), ch.GetClients)
		clientGroup.POST("/", handlerClientPrometheus(), ch.PostClient)
		clientGroup.PUT("/", handlerClientPrometheus(), ch.UpdateClient)
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
	l.Printf("Server is starting on %s\n", server.Addr)

	// Run the server on a new goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				// Normal interrupt operation, ignore
			} else {
				l.Fatalf("Server failed to start due to err: %v", err)
			}
		}
	}()

	// Block on this channel listening for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	l.Printf("Server is shutting down due to %+v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		l.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}
