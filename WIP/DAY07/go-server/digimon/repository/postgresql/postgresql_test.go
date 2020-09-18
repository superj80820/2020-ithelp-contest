package postgresql_test

import (
	"context"
	"testing"

	digimonPostgresqlRepo "go-server/digimon/repository/postgresql"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"go-server/domain"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDigimon := &domain.Digimon{
		ID:     "69770f2d-933e-474d-8357-a2f8a9c874df",
		Name:   "Customer",
		Status: "Good",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(mockDigimon.ID, mockDigimon.Name, mockDigimon.Status)

	query := "SELECT id FROM digimons WHERE id = ?"

	mock.ExpectQuery(query).WithArgs("69770f2d-933e-474d-8357-a2f8a9c874df").WillReturnRows(rows)
	d := digimonPostgresqlRepo.NewpostgresqlDigimonRepository(db)
	aDeit, _ := d.GetByID(context.TODO(), "69770f2d-933e-474d-8357-a2f8a9c874df")
	assert.Equal(t, mockDigimon, aDeit)
}

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDigimon := &domain.Digimon{
		ID:     "69770f2d-933e-474d-8357-a2f8a9c874df",
		Name:   "Customer",
		Status: "Good",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO digimons").
		WithArgs(mockDigimon.ID, mockDigimon.Name, mockDigimon.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	d := digimonPostgresqlRepo.NewpostgresqlDigimonRepository(db)

	tx, _ := db.Begin()
	d.Store(context.TODO(), mockDigimon)
	tx.Commit()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
