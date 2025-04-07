package server

import (
	"embed"
	"html/template"

	"github.com/mark3labs/mcp-go/server"

	"github.com/TuanKiri/weather-mcp-server/internal/server/services/core"
	"github.com/TuanKiri/weather-mcp-server/internal/server/tools"
	"github.com/TuanKiri/weather-mcp-server/pkg/weatherapi"
)

//go:embed view
var templates embed.FS

func Run(cfg *Config) error {
	tmpl, err := template.ParseFS(templates, "view/*.html")
	if err != nil {
		return err
	}

	wApi := weatherapi.New(cfg.WeatherAPIKey, cfg.WeatherAPITimeout)

	svc := core.New(tmpl, wApi)

	s := server.NewMCPServer(
		"Weather Server",
		"1.0.0",
		server.WithLogging(),
	)

	serverTools := []tools.ToolFunc{
		tools.CurrentWeather,
	}

	for _, tool := range serverTools {
		s.AddTool(tool(svc))
	}

	return server.ServeStdio(s)
}
