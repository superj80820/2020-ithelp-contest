package usecase

import (
	"context"

	"go-server/domain"
)

type dietUsecase struct {
	dietRepo domain.DietRepository
}

// NewDietUsecase ...
func NewDietUsecase(dietRepo domain.DietRepository) domain.DietUsecase {
	return &dietUsecase{
		dietRepo,
	}
}

func (du *dietUsecase) GetByID(ctx context.Context, id string) (*domain.Diet, error) {
	return nil, nil
}

func (du *dietUsecase) Store(ctx context.Context, d *domain.Diet) error {
	return nil
}
