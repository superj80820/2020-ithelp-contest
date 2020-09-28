package usecase

import (
	"context"

	"go-server/domain"

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

func (wu *weatherUsecase) GetStreamByLocation(ctx context.Context, location string) (domain.StreamWeather, error) {
	client, err := wu.weatherRepo.GetStreamByLocation(ctx, location)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return client, nil
}
