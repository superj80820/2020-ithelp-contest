package domain

import "context"

// Digimon ...
type Digimon struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// DigimonRepository ...
type DigimonRepository interface {
	GetByID(ctx context.Context, id string) (*Digimon, error)
	Store(ctx context.Context, d *Digimon) error
	UpdateStatus(ctx context.Context, d *Digimon) error
}

// DigimonUsecase ..
type DigimonUsecase interface {
	GetByID(ctx context.Context, id string) (*Digimon, error)
	Store(ctx context.Context, d *Digimon) error
	UpdateStatus(ctx context.Context, d *Digimon) error
}
