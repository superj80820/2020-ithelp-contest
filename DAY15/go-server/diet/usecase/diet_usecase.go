package usecase

import (
	"context"

	"go-server/domain"

	"github.com/sirupsen/logrus"
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
	aDiet, err := du.dietRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return aDiet, nil
}

func (du *dietUsecase) Store(ctx context.Context, d *domain.Diet) error {
	if err := du.dietRepo.Store(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
