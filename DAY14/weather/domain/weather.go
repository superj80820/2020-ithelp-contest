package domain

import "context"

// WeatherF ...
type WeatherEnum int32

const (
	// SUNNY ...
	SUNNY WeatherEnum = 0
	// CLOUDY ...
	CLOUDY WeatherEnum = 1
)

// Weather ...
type Weather struct {
	Location string
	Weather  WeatherEnum
}

// WeatherRepository ...
type WeatherRepository interface {
	GetByLocation(ctx context.Context, location string) (*Weather, error)
}

// WeatherUsecase ..
type WeatherUsecase interface {
	GetByLocation(ctx context.Context, location string) (*Weather, error)
}
