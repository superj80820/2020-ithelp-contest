package usecase

import (
	"context"

	"go-server/domain"

	"errors"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type digimonUsecase struct {
	digimonRepo domain.DigimonRepository
}

// NewDigimonUsecase ...
func NewDigimonUsecase(digimonRepo domain.DigimonRepository) domain.DigimonUsecase {
	return &digimonUsecase{
		digimonRepo: digimonRepo,
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
	if d.ID == "" {
		d.ID = uuid.Must(uuid.NewV4()).String()
	}
	if d.Status == "" {
		d.Status = "good"
	}
	if err := du.digimonRepo.Store(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (du *digimonUsecase) UpdateStatus(ctx context.Context, d *domain.Digimon) error {
	if d.Status == "" {
		err := errors.New("Status is blank")
		logrus.Error(err)
		return err
	}

	if err := du.digimonRepo.UpdateStatus(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
