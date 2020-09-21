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
	row := p.db.QueryRow("SELECT id, name, status FROM digimons WHERE id =$1", id)
	d := &domain.Digimon{}
	if err := row.Scan(&d.ID, &d.Name, &d.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return d, nil
}

func (p *postgresqlDigimonRepository) Store(ctx context.Context, d *domain.Digimon) error {
	_, err := p.db.Exec(
		"INSERT INTO digimons (id, name, status) VALUES ($1, $2, $3)",
		d.ID, d.Name, d.Status,
	)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (p *postgresqlDigimonRepository) UpdateStatus(ctx context.Context, d *domain.Digimon) error {
	_, err := p.db.Exec(
		"UPDATE digimons SET status=$1 WHERE id=$2",
		d.Status, d.ID,
	)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
