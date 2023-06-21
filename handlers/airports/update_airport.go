package airports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/models"
)

type UpdateAirportRequestBody struct {
	Name     string `json:"name" example:"John F. Kennedy International Airport" format:"string" swaggo:"true,Name,Name of the Airport"`
	CityID   uint   `json:"city_id" example:"1" format:"integer" swaggo:"true,CityID,The ID of the City"`
	IATACode string `json:"iata_code" example:"JFK" format:"string" swaggo:"true,IATACode,IATA code of the Airport"`
}

// @Summary Update airport
// @Description Update information of an airport given its ID
// @Tags Airports
// @Accept  json
// @Produce  json
// @Param id path int true "Airport ID"
// @Param body body UpdateAirportRequestBody true "Body with updated information"
// @Success 200 {object} models.Airport "Successfully updated airport information"
// @Failure 400 {object} error "Bad Request - Invalid or missing parameters"
// @Failure 404 {object} error "Not Found"
// @Router /airports/{id} [put]
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
