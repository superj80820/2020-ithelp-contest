package main

import (
	"net"
	_weatherHandlerGRPCDelivery "weather/weather/delivery/grpc"
	_weatherRepo "weather/weather/repository/fake"
	_weatherUsecase "weather/weather/usecase"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Fatal error config file: %v\n", err)
	}
}

func main() {
	logrus.Info("GRPC server started")

	grpcPort := viper.GetString("GRPC_PORT")

	weatherRepo := _weatherRepo.NewFakeWeatherRepository()

	weatherUsecase := _weatherUsecase.NewWeatherUsecase(weatherRepo)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		logrus.Fatal(err)
	}
	s := grpc.NewServer()

	_weatherHandlerGRPCDelivery.NewWeatherHandler(s, weatherUsecase)

	logrus.Fatal(s.Serve(lis))
}
