package main

import (
	"github.com/22Fariz22/binance-api/config"
	"github.com/22Fariz22/binance-api/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app := app.NewApp(cfg)
	app.Run()
}
