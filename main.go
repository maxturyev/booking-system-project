package main

import (
	"log"

	"github.com/maxturyev/booking-system-project/configs"
	"github.com/maxturyev/booking-system-project/db"
)

func main() {
	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := configs.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := configs.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Run the server
	cfg.Run()

	// Connect to system databases
	db.NewConnection()
}
