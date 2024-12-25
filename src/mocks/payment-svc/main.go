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
	pb "github.com/maxturyev/booking-system-project/mocks/grpc"
	"github.com/maxturyev/booking-system-project/payment-svc/common"
	"github.com/maxturyev/booking-system-project/payment-svc/db"
	grpcserver "github.com/maxturyev/booking-system-project/payment-svc/grpc-server"
	"github.com/maxturyev/booking-system-project/payment-svc/handlers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests.",
		},
		[]string{"method"},
	)
)

func handlerPaymentGetPrometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		elapsed := time.Since(start).Seconds()
		requestsTotal.WithLabelValues("GET").Inc()
		requestDuration.WithLabelValues("GET").Observe(elapsed)
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Обработка Прометея
	prometheus.MustRegister(requestsTotal, requestDuration)

	start := time.Now()
	elapsed := time.Since(start).Seconds()
	requestsTotal.WithLabelValues("GET").Inc()
	requestDuration.WithLabelValues("GET").Observe(elapsed)

	http.Handle("/metrics", promhttp.Handler())

	// Generate http server config
	cfg := common.NewConfig()

	// Create logger
	l := log.New(os.Stdout, "payment-svc\t", log.LstdFlags)

	// Connect to database
	hotelDb := db.ConnectDB()

	go func() {
		// Creating grpc-server
		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			l.Fatalf("Error starting server: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterHotelServiceServer(grpcServer, &grpcserver.HotelServer{DB: hotelDb})
		l.Println("Grpc server started successfully")
		if err := grpcServer.Serve(lis); err != nil {
			l.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	// Create router and define routes and return that router
	router := gin.Default()

	router.GET("/metrics", prometheusHandler())

	onlyH := handlers.NewPayments(l, hotelDb)
	// metricsH := handlers.NewPayments(l, hotelDb)
	// metricsGroup := router.Group("/metrics")
	// {
	// 	// paymentGroup.GET("/metrics", promhttp.Handler(), onlyH.ReturnError)
	// 	metricsGroup.GET("/", metricsH.DoPrometeus)
	// 	// paymentGroup.GET("/:id", validateNumericID(), onlyH.GetHotelByID)
	// 	// paymentGroup.POST("/", onlyH.PostHotel)
	// }
	paymentGroup := router.Group("/payment")
	{
		// paymentGroup.GET("/metrics", promhttp.Handler(), onlyH.ReturnError)
		paymentGroup.GET("/", handlerPaymentGetPrometheus(), onlyH.ReturnError)
		// paymentGroup.GET("/:id", validateNumericID(), onlyH.GetHotelByID)
		// paymentGroup.POST("/", onlyH.PostHotel)
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
				log.Fatalf("Server failed to start due to err: %v", err)
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
