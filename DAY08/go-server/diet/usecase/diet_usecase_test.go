package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	ucase "go-server/diet/usecase"
	"go-server/domain"
	"go-server/domain/mocks"
)

func TestGetByID(t *testing.T) {
	mockDietRepo := new(mocks.DietRepository)
	mockDiet := domain.Diet{
		ID:     "e9addf2d-8739-427a-8b30-2383b9b045b1",
		UserID: "ab18b1ba-48e1-48cf-88b5-48782874aa05",
		Name:   "Giuseppe",
	}

	t.Run("Success", func(t *testing.T) {
		mockDietRepo.
			On("GetByID", mock.Anything, mock.MatchedBy(func(value string) bool { return value == "e9addf2d-8739-427a-8b30-2383b9b045b1" })).
			Return(&mockDiet, nil).Once()

		u := ucase.NewDietUsecase(mockDietRepo)
		aDigimon, err := u.GetByID(context.TODO(), "e9addf2d-8739-427a-8b30-2383b9b045b1")

		assert.NoError(t, err)
		assert.NotNil(t, aDigimon)

		mockDietRepo.AssertExpectations(t)
	})
	t.Run("Fail", func(t *testing.T) {
		mockDietRepo.
			On("GetByID", mock.Anything, mock.MatchedBy(func(value string) bool { return value == "e9addf2d-8739-427a-8b30-2383b9b045b1" })).
			Return(nil, errors.New("Get error")).Once()

		u := ucase.NewDietUsecase(mockDietRepo)
		aDigimon, err := u.GetByID(context.TODO(), "e9addf2d-8739-427a-8b30-2383b9b045b1")

		assert.Error(t, err)
		assert.Nil(t, aDigimon)

		mockDietRepo.AssertExpectations(t)
	})
}

func TestStore(t *testing.T) {
	mockDietRepo := new(mocks.DietRepository)
	mockDiet := domain.Diet{
		ID:     "e9addf2d-8739-427a-8b30-2383b9b045b1",
		UserID: "ab18b1ba-48e1-48cf-88b5-48782874aa05",
		Name:   "Giuseppe",
	}

	t.Run("Success", func(t *testing.T) {
		mockDietRepo.
			On("Store", mock.Anything, mock.MatchedBy(func(d *domain.Diet) bool { return d == &mockDiet })).
			Return(nil).Once()

		u := ucase.NewDietUsecase(mockDietRepo)
		err := u.Store(context.TODO(), &mockDiet)

		assert.NoError(t, err)

		mockDietRepo.AssertExpectations(t)
	})
	t.Run("Fail", func(t *testing.T) {
		mockDietRepo.
			On("Store", mock.Anything, mock.MatchedBy(func(d *domain.Diet) bool { return d == &mockDiet })).
			Return(errors.New("Get error")).Once()

		u := ucase.NewDietUsecase(mockDietRepo)
		err := u.Store(context.TODO(), &mockDiet)

		assert.Error(t, err)

		mockDietRepo.AssertExpectations(t)
	})
}
