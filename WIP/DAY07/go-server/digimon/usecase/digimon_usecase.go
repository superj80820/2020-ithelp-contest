package usecase

import (
	"context"

	"go-server/domain"

	"github.com/sirupsen/logrus"
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
	aDigimon, err := du.digimonRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return aDigimon, nil
}

func (du *digimonUsecase) Store(ctx context.Context, d *domain.Digimon) error {
	if err := du.digimonRepo.Store(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
