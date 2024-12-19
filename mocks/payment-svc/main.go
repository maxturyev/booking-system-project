package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/payment-svc/common"
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
	// Generate http server config
	cfg := common.NewConfig()

	fmt.Println(cfg.Server.Host, cfg.Server.Port)

	// Create logger
	// l := log.New(os.Stdout, "payment-api", log.LstdFlags)

	// // Connect to database
	// hotelDb := db.ConnectDB()

	// go func() {
	// 	// Creating grpc-server
	// 	lis, err := net.Listen("tcp", ":50051")
	// 	if err != nil {
	// 		log.Fatalf("Error starting server: %v", err)
	// 	}

	// 	grpcServer := grpc.NewServer()
	// 	pb.RegisterHotelServiceServer(grpcServer, &grpcserver.HotelServer{DB: hotelDb})
	// 	log.Println("Grpc server started successfully")
	// 	if err := grpcServer.Serve(lis); err != nil {
	// 		log.Fatalf("Ошибка запуска сервера: %v", err)
	// 	}
	// }()

	// // Create router and define routes and return that router
	// router := gin.Default()

	// // Create handlers
	// hh := handlers.NewHotels(l, hotelDb)
	// hth := handlers.NewHoteliers(l, hotelDb)

	// // Handle requests for hotel
	// hotelGroup := router.Group("/hotel")
	// {
	// 	hotelGroup.GET("/", hh.GetHotels)
	// 	hotelGroup.GET("/:id", validateNumericID(), hh.GetHotelByID)
	// 	hotelGroup.POST("/", hh.PostHotel)
	// }

	// // Handle requests for hotelier
	// hotelierGroup := router.Group("/hotelier")
	// {
	// 	hotelierGroup.GET("/", hth.GetHoteliers)
	// 	hotelierGroup.POST("/", hth.PostHotelier)
	// }

	// // Set up a channel to listen to for interrupt signals
	// var runChan = make(chan os.Signal, 1)

	// // Set up a context to allow for graceful server shutdowns in the event
	// // of an OS interrupt (defers the cancel just in case)
	// ctx, cancel := context.WithTimeout(
	// 	context.Background(),
	// 	cfg.Server.Timeout.Server,
	// )
	// defer cancel()

	// // Define server options
	// server := &http.Server{
	// 	Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
	// 	Handler:      router,
	// 	ReadTimeout:  cfg.Server.Timeout.Read * time.Second,
	// 	WriteTimeout: cfg.Server.Timeout.Write * time.Second,
	// 	IdleTimeout:  cfg.Server.Timeout.Idle * time.Second,
	// }

	// // Handle ctrl+c/ctrl+x interrupt
	// signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	// // Alert the user that the server is starting
	// log.Printf("Server is starting on %s\n", server.Addr)

	// // Run the server on a new goroutine
	// go func() {
	// 	if err := server.ListenAndServe(); err != nil {
	// 		if errors.Is(err, http.ErrServerClosed) {
	// 			// Normal interrupt operation, ignore
	// 		} else {
	// 			log.Fatalf("Server failed to start due to err: %v", err)
	// 		}
	// 	}
	// }()

	// // Block on this channel listening for those previously defined syscalls assign
	// // to variable so we can let the user know why the server is shutting down
	// interrupt := <-runChan

	// // If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// // while alerting the user
	// log.Printf("Server is shutting down due to %+v\n", interrupt)
	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	// }

}
