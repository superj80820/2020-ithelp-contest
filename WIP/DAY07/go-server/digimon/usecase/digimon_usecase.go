package usecase

import (
	"context"

	"go-server/domain"
)

type digimonUsecase struct {
	digimonRepo domain.DigimonRepository
}

// NewDigimonUsecase ...
func NewDigimonUsecase(digimonRepo domain.DigimonRepository) domain.DigimonUsecase {
	return &digimonUsecase{
		digimonRepo,
	}
}

func (du *digimonUsecase) GetByID(ctx context.Context, id string) (*domain.Digimon, error) {
	return nil, nil
}

func (du *digimonUsecase) Store(ctx context.Context, d *domain.Digimon) error {
	return nil
}
