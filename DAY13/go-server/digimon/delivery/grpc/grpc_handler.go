package grpc

import (
	"context"

	"go-server/domain"
	pb "go-server/gen/grpc-gateway/gen"

	"github.com/sirupsen/logrus"
	grpcLib "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DigimonHandler ...
type DigimonHandler struct {
	DigimonUsecase domain.DigimonUsecase
	DietUsecase    domain.DietUsecase
	pb.UnimplementedDigimonServer
}

// NewDigimonHandler ...
func NewDigimonHandler(s *grpcLib.Server, digimonUsecase domain.DigimonUsecase, dietUsecase domain.DietUsecase) {
	handler := &DigimonHandler{
		DigimonUsecase: digimonUsecase,
		DietUsecase:    dietUsecase,
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

// Query ...
func (d *DigimonHandler) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	anDigimon, err := d.DigimonUsecase.GetByID(ctx, req.GetId())
	if err != nil {
		logrus.Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error. Query digimon error")
	}

	return &pb.QueryResponse{
		Id:     anDigimon.ID,
		Name:   anDigimon.Name,
		Status: anDigimon.Status,
	}, nil
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
