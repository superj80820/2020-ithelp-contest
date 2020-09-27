package http

import (
	"go-server/domain"

	swagger "go-server/go"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// DigimonHandler ...
type DigimonHandler struct {
	DigimonUsecase domain.DigimonUsecase
	DietUsecase    domain.DietUsecase
}

// NewDigimonHandler ...
func NewDigimonHandler(e *gin.Engine, digimonUsecase domain.DigimonUsecase, dietUsecase domain.DietUsecase) {
	handler := &DigimonHandler{
		DigimonUsecase: digimonUsecase,
		DietUsecase:    dietUsecase,
	}

	e.GET("/api/v1/digimons/:digimonID", handler.GetDigimonByDigimonID)
	e.POST("/api/v1/digimons", handler.PostToCreateDigimon)
	e.POST("/api/v1/digimons/:digimonID/foster", handler.PostToFosterDigimon)
}

// GetDigimonByDigimonID ...
func (d *DigimonHandler) GetDigimonByDigimonID(c *gin.Context) {
	digimonID := c.Param("digimonID")

	anDigimon, err := d.DigimonUsecase.GetByID(c, digimonID)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Query digimon error",
		})
		return
	}

	c.JSON(200, &swagger.DigimonInfo{
		Id:     anDigimon.ID,
		Name:   anDigimon.Name,
		Status: anDigimon.Status,
	})
}

// PostToCreateDigimon ...
func (d *DigimonHandler) PostToCreateDigimon(c *gin.Context) {
	var body swagger.DigimonInfoRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Parsing failed",
		})
		return
	}
	aDigimon := domain.Digimon{
		Name: body.Name,
	}
	if err := d.DigimonUsecase.Store(c, &aDigimon); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}

	c.JSON(200, swagger.DigimonInfo{
		Id:     aDigimon.ID,
		Name:   aDigimon.Name,
		Status: aDigimon.Status,
	})
}

// PostToFosterDigimon ...
func (d *DigimonHandler) PostToFosterDigimon(c *gin.Context) {
	digimonID := c.Param("digimonID")

	var body swagger.FosterRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Parsing failed",
		})
		return
	}

	if err := d.DietUsecase.Store(c, &domain.Diet{
		UserID: digimonID,
		Name:   body.Food.Name,
	}); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}

	if err := d.DigimonUsecase.UpdateStatus(c, &domain.Digimon{
		ID:     digimonID,
		Status: "good",
	}); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}
	c.JSON(204, nil)
}
