package grpc

import (
	"context"
	"errors"
	"io"

	"weather/domain"
	pb "weather/gen"

	"github.com/sirupsen/logrus"
	grpcLib "google.golang.org/grpc"
)

// WeatherHandler ...
type WeatherHandler struct {
	WeatherUsecase domain.WeatherUsecase
	pb.UnimplementedWeatherServer
}

// NewWeatherHandler ...
func NewWeatherHandler(s *grpcLib.Server, weatherUsecase domain.WeatherUsecase) {
	handler := &WeatherHandler{
		WeatherUsecase: weatherUsecase,
	}

	pb.RegisterWeatherServer(s, handler)
}

func mappingGRPCWeatherEnum(weather domain.WeatherEnum) (pb.QueryResponse_Weather, error) {
	switch weather {
	case domain.SUNNY:
		return pb.QueryResponse_SUNNY, nil
	case domain.CLOUDY:
		return pb.QueryResponse_CLOUDY, nil
	default:
		return pb.QueryResponse_SUNNY, errors.New("This weather does not exist")
	}
}

// Query ...
func (w *WeatherHandler) Query(srv pb.Weather_QueryServer) error {
	for {
		msg, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			logrus.Error(err)
			return err
		}

		aWeather, err := w.WeatherUsecase.GetByLocation(context.Background(), msg.GetLocation())
		if err != nil {
			logrus.Error(err)
		}

		gRPCWriterEnum, err := mappingGRPCWeatherEnum(aWeather.Weather)
		if err != nil {
			logrus.Error(err)
		}

		srv.Send(&pb.QueryResponse{
			Location: aWeather.Location,
			Weather:  gRPCWriterEnum,
		})
	}
}
