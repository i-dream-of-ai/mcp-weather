package main

import (
	"log"
	"os"
	"time"

	"github.com/TuanKiri/weather-mcp-server/internal/server"
)

func main() {
	cfg := server.Config{
		WeatherAPIKey:     os.Getenv("WEATHER_API_KEY"),
		WeatherAPITimeout: 1 * time.Second,
	}

	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(&cfg); err != nil {
		log.Fatal(err)
	}
}
