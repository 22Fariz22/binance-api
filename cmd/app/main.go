package app

import (
	"github.com/22Fariz22/binance-api/internal/app"
	"github.com/22Fariz22/binance-api/internal/config"
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
