package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/hotel-svc/common"
	"github.com/maxturyev/booking-system-project/hotel-svc/db"
	grpcserver "github.com/maxturyev/booking-system-project/hotel-svc/grpc-server"
	"github.com/maxturyev/booking-system-project/hotel-svc/handlers"
	pb "github.com/maxturyev/booking-system-project/src/grpc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

var (
	requestsTotalHotel = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total_hotel",
			Help: "Total number of HTTP requests on hotel.",
		},
		[]string{"method"},
	)
	requestDurationHotel = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds_hotel",
			Help: "Duration of HTTP requests on hotel.",
		},
		[]string{"method"},
	)
	requestsTotalHotelier = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total_hotelier",
			Help: "Total number of HTTP requests on hotelier.",
		},
		[]string{"method"},
	)
	requestDurationHotelier = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds_hotelier",
			Help: "Duration of HTTP requests on hotelier.",
		},
		[]string{"method"},
	)
)

func handlerHotelPrometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		start := time.Now()
		elapsed := time.Since(start).Seconds()
		requestsTotalHotel.WithLabelValues(method).Inc()
		requestDurationHotel.WithLabelValues(method).Observe(elapsed)
	}
}

func handlerHotelierPrometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		start := time.Now()
		elapsed := time.Since(start).Seconds()
		requestsTotalHotelier.WithLabelValues(method).Inc()
		requestDurationHotelier.WithLabelValues(method).Observe(elapsed)
	}
}

func prometheusView() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

const PORT = ":50051"

func main() {
	// Обработка Прометея
	prometheus.MustRegister(requestsTotalHotel, requestDurationHotel, requestsTotalHotelier, requestDurationHotelier)

	// Generate httpServer config
	cfg := common.NewConfig()

	// Create logger
	l := log.New(os.Stdout, "hotel-svc\t", log.LstdFlags)

	// Connect to database
	hotelDb := db.ConnectDB()

	// Create grpc httpServer
	grpcServer := grpc.NewServer()

	// Run the http httpServer on a new goroutine
	go func() {
		lis, err := net.Listen("tcp", PORT)
		if err != nil {
			l.Fatalf("failed to create listener: %s", err)
		}

		pb.RegisterHotelServiceServer(grpcServer, &grpcserver.HotelServer{DB: hotelDb})

		l.Printf("Grpc httpServer started on port %s", PORT)
		if err := grpcServer.Serve(lis); err != nil {
			l.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	// Create router and define routes and return that router
	router := gin.Default()

	router.GET("/metrics", prometheusView())

	// Create handlers
	hh := handlers.NewHotels(l, hotelDb)
	hth := handlers.NewHoteliers(l, hotelDb)

	// Handle requests for hotel
	hotelGroup := router.Group("/hotel")
	{
		hotelGroup.GET("/", handlerHotelPrometheus(), hh.GetHotels)
		hotelGroup.GET("/:id", handlerHotelPrometheus(), handlers.ValidateNumericID(), hh.GetHotelByID)
		hotelGroup.POST("/", handlerHotelPrometheus(), hh.PostHotel)
		hotelGroup.POST("/media", handlerHotelPrometheus(), hh.HandleUploadImage)
	}

	// Handle requests for hotelier
	hotelierGroup := router.Group("/hotelier")
	{
		hotelierGroup.GET("/", handlerHotelierPrometheus(), hth.GetHoteliers)
		hotelierGroup.POST("/", handlerHotelierPrometheus(), hth.PostHotelier)
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
	l.Printf("Server is starting on %s\n", httpServer.Addr)

	// Run the http httpServer on a new goroutine
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				// Normal interrupt operation, ignore
			} else {
				l.Fatalf("Server failed to start due to err: %v", err)
			}
		}
	}()

	// Block on this channel listening for those previously defined syscalls assign
	// to variable so we can let the user know why the httpServer is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the httpServer
	// while alerting the user
	l.Printf("Server is shutting down due to %+v\n", interrupt)
	if err := httpServer.Shutdown(ctx); err != nil {
		l.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}
