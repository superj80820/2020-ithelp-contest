package grpc

import (
	"context"

	"go-server/domain"
	pb "go-server/gen/weather/gen"

	"github.com/sirupsen/logrus"
)

type grpcWeatherRepository struct {
	weatherGRPC pb.WeatherClient
}

type weatherClient struct {
	client pb.Weather_QueryClient
}

func (wc *weatherClient) Send(w *domain.Weather) error {
	if err := wc.client.Send(&pb.QueryRequest{
		Location: w.Location,
	}); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (wc *weatherClient) Recv() (*domain.Weather, error) {
	weather, err := wc.client.Recv()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &domain.Weather{
		Location: weather.GetLocation(),
		Weather:  string(weather.GetWeather()),
	}, nil
}

// NewgrpcWeatherRepository ...
func NewgrpcWeatherRepository(weatherGRPC pb.WeatherClient) domain.WeatherRepository {
	return &grpcWeatherRepository{
		weatherGRPC,
	}
}

func (g *grpcWeatherRepository) GetStreamByLocation(ctx context.Context, location string) (domain.StreamWeather, error) {
	clinet, err := g.weatherGRPC.Query(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &weatherClient{
		clinet,
	}, nil
}
