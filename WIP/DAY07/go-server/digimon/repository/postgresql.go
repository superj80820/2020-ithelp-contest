package postgresql

import (
	"context"
	"database/sql"

	"go-server/domain"

	_ "github.com/lib/pq"
)

type postgresqlDigimonRepository struct {
	db *sql.DB
}

// NewpostgresqlDigimonRepository ...
func NewpostgresqlDigimonRepository(db *sql.DB) domain.DigimonRepository {
	return &postgresqlDigimonRepository{db}
}

func (p *postgresqlDigimonRepository) GetByID(ctx context.Context, id string) (*domain.Digimon, error) {
	return nil, nil
}

func (p *postgresqlDigimonRepository) Store(ctx context.Context, d *domain.Digimon) error {
	return nil
}
