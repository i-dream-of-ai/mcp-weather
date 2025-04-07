package services

import "context"

type Services interface {
	Weather() WeatherService
}

type WeatherService interface {
	Current(ctx context.Context, city string) (string, error)
}
