package domain

import "context"

// Diet ...
type Diet struct {
	ID     string
	UserID string
	Name   string
}

// DietRepository ...
type DietRepository interface {
	GetByID(ctx context.Context, id string) (*Diet, error)
	Store(ctx context.Context, d *Diet) error
}

// DietUsecase ..
type DietUsecase interface {
	GetByID(ctx context.Context, id string) (*Diet, error)
	Store(ctx context.Context, d *Diet) error
}
