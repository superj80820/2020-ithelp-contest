package postgresql

import (
	"context"
	"database/sql"

	"go-server/domain"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type postgresqlDietRepository struct {
	db *sql.DB
}

// NewPostgresqlDietRepository ...
func NewPostgresqlDietRepository(db *sql.DB) domain.DietRepository {
	return &postgresqlDietRepository{db}
}

func (p *postgresqlDietRepository) GetByID(ctx context.Context, id string) (*domain.Diet, error) {
	row := p.db.QueryRow("SELECT id FROM diets WHERE id = ?", id)
	d := &domain.Diet{}
	if err := row.Scan(&d.ID, &d.UserID, &d.Name); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return d, nil
}

func (p *postgresqlDietRepository) Store(ctx context.Context, d *domain.Diet) error {
	return nil
}
