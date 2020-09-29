package grpc

import (
	"context"
	"time"

	"go-server/domain"
	pb "go-server/gen"

	"github.com/sirupsen/logrus"
	grpcLib "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DigimonHandler ...
type DigimonHandler struct {
	DigimonUsecase domain.DigimonUsecase
	DietUsecase    domain.DietUsecase
	WeatherUsecase domain.WeatherUsecase
	pb.UnimplementedDigimonServer
}

// NewDigimonHandler ...
func NewDigimonHandler(s *grpcLib.Server, digimonUsecase domain.DigimonUsecase, dietUsecase domain.DietUsecase, weatherUsecase domain.WeatherUsecase) {
	handler := &DigimonHandler{
		DigimonUsecase: digimonUsecase,
		DietUsecase:    dietUsecase,
		WeatherUsecase: weatherUsecase,
	}

	pb.RegisterDigimonServer(s, handler)
}

// Create ...
func (d *DigimonHandler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	aDigimon := domain.Digimon{
		Name: req.GetName(),
	}
	if err := d.DigimonUsecase.Store(ctx, &aDigimon); err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Store failed")
	}

	return &pb.CreateResponse{
		Id:     aDigimon.ID,
		Name:   aDigimon.Name,
		Status: aDigimon.Status,
	}, nil
}

// QueryStream ...
func (d *DigimonHandler) QueryStream(req *pb.QueryRequest, srv pb.Digimon_QueryStreamServer) error {
	weatherClient, err := d.WeatherUsecase.GetStreamByLocation(context.Background(), "A")
	if err != nil {
		logrus.Error(err)
		return err
	}

	for {
		if err := weatherClient.Send(&domain.Weather{
			Location: "A",
		}); err != nil {
			logrus.Error(err)
			return err
		}

		time.Sleep(time.Duration(5) * time.Second)

		aWeather, err := weatherClient.Recv()
		if err != nil {
			logrus.Error(err)
			return err
		}

		aDigimon, err := d.DigimonUsecase.GetByID(context.Background(), req.GetId())
		if err != nil {
			logrus.Error(err)
			return err
		}

		srv.Send(&pb.QueryResponse{
			Id:       aDigimon.ID,
			Name:     aDigimon.Name,
			Status:   aDigimon.Status,
			Location: aWeather.Location,
			Weather:  aWeather.Weather,
		})
	}
}

// Foster ...
func (d *DigimonHandler) Foster(ctx context.Context, req *pb.FosterRequest) (*pb.FosterResponse, error) {
	if err := d.DietUsecase.Store(ctx, &domain.Diet{
		UserID: req.GetId(),
		Name:   req.GetFood().GetName(),
	}); err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Store failed")
	}

	if err := d.DigimonUsecase.UpdateStatus(ctx, &domain.Digimon{
		ID:     req.GetId(),
		Status: "good",
	}); err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Store failed")
	}

	return &pb.FosterResponse{}, nil
}
