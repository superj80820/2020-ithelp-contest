package usecase

import (
	"context"

	"weather/domain"

	"github.com/sirupsen/logrus"
)

type weatherUsecase struct {
	weatherRepo domain.WeatherRepository
}

// NewWeatherUsecase ...
func NewWeatherUsecase(weatherRepo domain.WeatherRepository) domain.WeatherUsecase {
	return &weatherUsecase{
		weatherRepo: weatherRepo,
	}
}

func (w *weatherUsecase) GetByLocation(ctx context.Context, location string) (*domain.Weather, error) {
	aWeather, err := w.weatherRepo.GetByLocation(ctx, location)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return aWeather, nil
}
