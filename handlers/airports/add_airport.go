package airports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type AddAirportRequestBody struct {
	Name     string `json:"name" example:"John F. Kennedy International Airport" format:"string" swaggo:"true,Name,Name of the Airport"`
	CityID   uint   `json:"city_id" example:"1" format:"integer" swaggo:"true,CityID,The ID of the City"`
	IATACode string `json:"iata_code" example:"JFK" format:"string" swaggo:"true,IATACode,IATA code of the Airport"`
}

// @Summary Add a new airport
// @Description Add a new airport to the system
// @Tags Airports
// @Accept  json
// @Produce  json
// @Param body body AddAirportRequestBody true "Airport details"
// @Success 201 {object} models.Airport
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Not Found"
// @Router /airports [post]
func (h handler) AddAirport(c *gin.Context) {
	body := AddAirportRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var airport models.Airport

	airport.CityID = body.CityID
	airport.Name = body.Name
	airport.IATACode = body.IATACode

	if result := h.DB.Create(&airport); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &airport)
}
