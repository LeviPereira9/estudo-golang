package main

import (
	"fmt"
	"log"
	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"

)

// config -> db -> rotuer -> run server
func main(){

	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Config error")
	}

	client, database, err := db.Connect(cfg)

	if err != nil {
		log.Fatalf("Connection error")
	}

	defer func(){

		if err := db.Disconnect(client); err != nil {
			log.Printf("mongo disconnect error: %v", err)
		}
		
	}()

	router := server.NewRouter(database)

	address := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(address); err != nil {
		log.Fatalf("server failed")
	}
	
	
}
