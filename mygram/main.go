package main

import (
	"fmt"
	"log"
	"mygram/config"
	"mygram/database"
	"mygram/router"
)

func main() {
	// load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// initialize database
	db, err := database.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// initialize router
	r := router.InitRouter()

	// start server
	addr := fmt.Sprintf(":%d", cfg.Port)
	err = r.Run(addr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
