package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"errors"
	ucase "go-server/digimon/usecase"
	"go-server/domain"
	"go-server/domain/mocks"
)

func TestGetByID(t *testing.T) {
	mockDigimonRepo := new(mocks.DigimonRepository)
	mockDigimon := domain.Digimon{
		ID:     "2e72c27e-0feb-44e1-89ef-3d58fd30a1b3",
		Name:   "Metrics",
		Status: "Good",
	}

	t.Run("Success", func(t *testing.T) {
		mockDigimonRepo.
			On("GetByID", mock.Anything, mock.MatchedBy(func(value string) bool { return value == "e9addf2d-8739-427a-8b30-2383b9b045b1" })).
			Return(&mockDigimon, nil).Once()

		u := ucase.NewDigimonUsecase(mockDigimonRepo)
		aDigimon, err := u.GetByID(context.TODO(), "e9addf2d-8739-427a-8b30-2383b9b045b1")

		assert.NoError(t, err)
		assert.NotNil(t, aDigimon)

		mockDigimonRepo.AssertExpectations(t)
	})
	t.Run("Fail", func(t *testing.T) {
		mockDigimonRepo.
			On("GetByID", mock.Anything, mock.MatchedBy(func(value string) bool { return value == "e9addf2d-8739-427a-8b30-2383b9b045b1" })).
			Return(nil, errors.New("Get error")).Once()

		u := ucase.NewDigimonUsecase(mockDigimonRepo)
		aDigimon, err := u.GetByID(context.TODO(), "e9addf2d-8739-427a-8b30-2383b9b045b1")

		assert.Error(t, err)
		assert.Nil(t, aDigimon)

		mockDigimonRepo.AssertExpectations(t)
	})
}

func TestStore(t *testing.T) {
	mockDigimonRepo := new(mocks.DigimonRepository)
	mockDigimon := domain.Digimon{
		ID:     "2e72c27e-0feb-44e1-89ef-3d58fd30a1b3",
		Name:   "Metrics",
		Status: "Good",
	}

	t.Run("Success", func(t *testing.T) {
		mockDigimonRepo.
			On("Store", mock.Anything, mock.MatchedBy(func(d *domain.Digimon) bool { return d == &mockDigimon })).
			Return(nil).Once()

		u := ucase.NewDigimonUsecase(mockDigimonRepo)
		err := u.Store(context.TODO(), &mockDigimon)

		assert.NoError(t, err)

		mockDigimonRepo.AssertExpectations(t)
	})
	t.Run("Fail", func(t *testing.T) {
		mockDigimonRepo.
			On("Store", mock.Anything, mock.MatchedBy(func(d *domain.Digimon) bool { return d == &mockDigimon })).
			Return(errors.New("Get error")).Once()

		u := ucase.NewDigimonUsecase(mockDigimonRepo)
		err := u.Store(context.TODO(), &mockDigimon)

		assert.Error(t, err)

		mockDigimonRepo.AssertExpectations(t)
	})
}
