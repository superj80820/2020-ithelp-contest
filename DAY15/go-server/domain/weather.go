package domain

import "context"

// Weather ...
type Weather struct {
	Location string
	Weather  string
}

// StreamWeather ...
type StreamWeather interface {
	Send(*Weather) error
	Recv() (*Weather, error)
}

// WeatherRepository ...
type WeatherRepository interface {
	GetStreamByLocation(ctx context.Context, location string) (StreamWeather, error)
}

// WeatherUsecase ..
type WeatherUsecase interface {
	GetStreamByLocation(ctx context.Context, location string) (StreamWeather, error)
}
