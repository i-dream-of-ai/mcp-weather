package core

import (
	"html/template"

	"github.com/TuanKiri/weather-mcp-server/internal/server/services"
	"github.com/TuanKiri/weather-mcp-server/pkg/weatherapi"
)

type CoreServices struct {
	renderer   *template.Template
	weatherAPI *weatherapi.WeatherAPI

	weatherService *WeatherService
}

func New(renderer *template.Template, weatherAPI *weatherapi.WeatherAPI) *CoreServices {
	return &CoreServices{
		renderer:   renderer,
		weatherAPI: weatherAPI,
	}
}

func (cs *CoreServices) Weather() services.WeatherService {
	if cs.weatherService == nil {
		cs.weatherService = &WeatherService{CoreServices: cs}
	}

	return cs.weatherService
}
