package postgresql

import (
	"context"
	"database/sql"

	"go-server/domain"

	"github.com/sirupsen/logrus"
)

type postgresqlDigimonRepository struct {
	db *sql.DB
}

// NewpostgresqlDigimonRepository ...
func NewpostgresqlDigimonRepository(db *sql.DB) domain.DigimonRepository {
	return &postgresqlDigimonRepository{db}
}

func (p *postgresqlDigimonRepository) GetByID(ctx context.Context, id string) (*domain.Digimon, error) {
	row := p.db.QueryRow("SELECT id FROM digimons WHERE id = ?", id)
	d := &domain.Digimon{}
	if err := row.Scan(&d.ID, &d.Name, &d.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return d, nil
}

func (p *postgresqlDigimonRepository) Store(ctx context.Context, d *domain.Digimon) error {
	_, err := p.db.Exec(
		"INSERT INTO digimons (id, name, status) VALUES (?, ?, ?)",
		d.ID, d.Name, d.Status,
	)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
