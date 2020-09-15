package domain

import "context"

// Digimon ...
type Digimon struct {
	ID     string
	Name   string
	Status string
}

// DigimonRepository ...
type DigimonRepository interface {
	GetByID(ctx context.Context, id string) (*Digimon, error)
	Store(ctx context.Context, d *Digimon) error
}

// DigimonUsecase ..
type DigimonUsecase interface {
	GetByID(ctx context.Context, id string) (*Digimon, error)
	Store(ctx context.Context, d *Digimon) error
}
