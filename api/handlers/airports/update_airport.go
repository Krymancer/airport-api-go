package airports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type UpdateAirportRequestBody struct {
	Name     string `json:"name"`
	CityID   uint   `json:"city_id"`
	IATACode string `json:"iata_code"`
}

func (h handler) UpdateAirport(c *gin.Context) {
	id := c.Param("id")
	body := UpdateAirportRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var airport models.Airport

	if result := h.DB.First(&airport, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	airport.Name = body.Name
	airport.CityID = body.CityID
	airport.IATACode = body.IATACode

	h.DB.Save(&airport)

	c.JSON(http.StatusOK, &airport)
}
