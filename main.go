package main

import (
	"log"
	"resto-backend/config"
	"resto-backend/server"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server %:", err)
	}

	err = server.Start(config.PORT)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
