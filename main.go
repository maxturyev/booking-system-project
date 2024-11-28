
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maxturyev/booking-system-project/api/handlers"
)

func main() {
	l := log.New(os.Stdout, "hotel-api", log.LstdFlags)

	// create a hotel handler
	hh := handlers.NewHotels(l)
	ch := handlers.NewClient(l)

	// create a new serve mux and register the handler
	sm := http.NewServeMux()
	sm.Handle("/hotel/", hh)
	sm.Handle("/client/", ch)

	// create a new server
	s := http.Server{
		Addr:     ":9090", // TCP address
		Handler:  sm,      // Default handler
		ErrorLog: l,       // Logger
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		if err := s.ListenAndServe(); err != nil {
			l.Fatalf("Error starting server: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	l.Println("Server exited properly")
}

