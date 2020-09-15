package postgresql_test

import (
	"context"
	"testing"

	deitPostgresqlRepo "go-server/diet/repository/postgresql"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"go-server/domain"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDiet := &domain.Diet{
		ID:     "ad1de98f-c0cd-46c6-9349-eec80f4b1a12",
		UserID: "318f5799-e2f9-4ace-96e0-6658fe603d10",
		Name:   "Apple",
	}

	rows := sqlmock.NewRows([]string{"id", "userID", "name"}).
		AddRow(mockDiet.ID, mockDiet.UserID, mockDiet.Name)

	query := "SELECT id FROM diets WHERE id = ?"

	mock.ExpectQuery(query).WithArgs("ad1de98f-d5ec-4976-867e-531429a28cda").WillReturnRows(rows)
	d := deitPostgresqlRepo.NewPostgresqlDietRepository(db)
	aDeit, err := d.GetByID(context.TODO(), "ad1de98f-d5ec-4976-867e-531429a28cda")
	assert.Equal(t, mockDiet, aDeit)
}
