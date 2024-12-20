package main

import (
	"jual-beli-barang-bekas/config"
	"jual-beli-barang-bekas/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("Config file not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}
