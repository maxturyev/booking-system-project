package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maxturyev/booking-system-project/booking-svc/common"
	pb "github.com/maxturyev/booking-system-project/src/grpc"
	"google.golang.org/grpc"

	"github.com/maxturyev/booking-system-project/booking-svc/databases"
	"github.com/maxturyev/booking-system-project/booking-svc/handlers"
)

func main() {
	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := common.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := common.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create logger
	l := log.New(os.Stdout, "booking-svc", log.LstdFlags)

	// Connect to database
	db, err := databases.Init()
	if err != nil {
		panic(err)
	}
	//grpc client server connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("error: can not connection")
	}
	defer conn.Close()
	clientgrpc := pb.NewHotelServiceClient(conn)

	router := gin.Default()

	bh := handlers.NewBookings(l, db)
	ch := handlers.NewClients(l, db)

	bookingGroup := router.Group("/booking")
	{
		client := bookingGroup.Group("/client")
		{
			client.GET("/", bh.ListBookings)
			client.POST("/", bh.CreateBooking)
			client.PUT("/", bh.UpdateBooking)
		}
		bookingGroup.GET("/hotels", func(ctx *gin.Context) {
			ctxgrpc, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			log.Println("in grpc")
			stream, err := clientgrpc.GetHotels(ctxgrpc, &pb.GetHotelsRequest{})
			if err != nil {
				log.Fatal("error")
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
					log.Print("Ended")
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
					RoomsAvailable: int(res.Hotel.RoomAvaible),
					Price:          int(res.Hotel.Price),
					Address:        res.Hotel.Address,
				})
			}

			ctx.JSON(200, gin.H{
				"hotels": hotelList,
			})
		})
		bookingGroup.GET("/", ch.ListClients)
		bookingGroup.POST("/", ch.AddClient)
		bookingGroup.PUT("/", ch.UpdateClient)

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
			if err == http.ErrServerClosed {
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
