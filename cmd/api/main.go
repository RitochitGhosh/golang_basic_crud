package main

import (
	"fmt"
	"go-crud/internal/config"
	"go-crud/internal/db"
	"go-crud/internal/server"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error")
	}

	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("DB error: ", err)
	}

	defer func() {
		if err := db.Disconnect(client); err != nil {
			log.Printf("Mongo disconnect error: %v", err)
		}
	}()

	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Server failed")
	}

}
