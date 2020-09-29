package http_test

import (
	"bytes"
	"encoding/json"
	"go-server/domain"
	"go-server/domain/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	digimonHandlerHttpDelivery "go-server/digimon/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDigimonByDigimonID(t *testing.T) {
	mockDigimon := domain.Digimon{
		ID:     "8c862535-6de2-4da2-ad21-c853e4343bd7",
		Name:   "III",
		Status: "good",
	}
	mockDigimonMarshal, _ := json.Marshal(mockDigimon)
	mockDigimonCase := new(mocks.DigimonUsecase)

	mockDigimonCase.On("GetByID", mock.Anything, mockDigimon.ID).Return(&mockDigimon, nil)

	r := gin.Default()

	handler := digimonHandlerHttpDelivery.DigimonHandler{
		DigimonUsecase: mockDigimonCase,
	}

	r.GET("/api/v1/digimons/:digimonID", handler.GetDigimonByDigimonID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/digimons/8c862535-6de2-4da2-ad21-c853e4343bd7", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(mockDigimonMarshal), w.Body.String())
}

func TestPostToCreateDigimon(t *testing.T) {
	mockDigimonCase := new(mocks.DigimonUsecase)

	mockDigimonCase.
		On("Store", mock.Anything, mock.AnythingOfType("*domain.Digimon")).
		Return(nil).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Digimon)
			arg.ID = "e5d88876-e513-43e9-80ac-6b348d84d8b4"
			arg.Status = "good"
		})

	r := gin.Default()

	handler := digimonHandlerHttpDelivery.DigimonHandler{
		DigimonUsecase: mockDigimonCase,
	}

	r.POST("/api/v1/digimons", handler.PostToCreateDigimon)

	jsonString := []byte(`{"name":"Agumon"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/digimons", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":"e5d88876-e513-43e9-80ac-6b348d84d8b4","name":"Agumon","status":"good"}`, w.Body.String())
}

func TestPostToFosterDigimon(t *testing.T) {
	mockDigimonCase := new(mocks.DigimonUsecase)
	mockDietCase := new(mocks.DietUsecase)

	mockDietCase.
		On("Store", mock.Anything, mock.MatchedBy(func(value *domain.Diet) bool {
			return value.UserID == "178744e4-a218-45ac-adfa-9023a3bf9699" && value.Name == "apple"
		})).
		Return(nil)

	mockDigimonCase.
		On("UpdateStatus", mock.Anything, mock.Anything).
		Return(nil)

	r := gin.Default()

	handler := digimonHandlerHttpDelivery.DigimonHandler{
		DigimonUsecase: mockDigimonCase,
		DietUsecase:    mockDietCase,
	}

	r.POST("/api/v1/digimons/:digimonID/foster", handler.PostToFosterDigimon)

	jsonString := []byte(`{"food": {"name": "apple"}}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/digimons/178744e4-a218-45ac-adfa-9023a3bf9699/foster", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}
