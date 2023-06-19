package airports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type AddAirportRequestBody struct {
	Name     string      `json:"name"`
	City     models.City `json:"city"`
	IATACode string      `json:"iata_code"`
}

func (h handler) AddAirport(c *gin.Context) {
	body := AddAirportRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var airport models.Airport

	airport.City = body.City
	airport.Name = body.Name
	airport.IATACode = body.IATACode

	if result := h.DB.Create(&airport); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &airport)
}
