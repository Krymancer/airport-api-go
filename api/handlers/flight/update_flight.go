package flights

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type UpdateFlightBodyRequest struct {
	Origin        string               `json:"origin" binding:"required"`
	Destination   string               `json:"destination" binding:"required"`
	FlightClasses []models.FlightClass `json:"flight_classes" binding:"required"`
	Number        string               `json:"number" binding:"required"`
}

func (h handler) UpdateFlight(c *gin.Context) {
	id := c.Param("id")
	body := UpdateFlightBodyRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var flight models.Flight
	var origin models.Airport
	var destination models.Airport

	if result := h.DB.Where("number = ?", body.Number).First(&flight); result.Error == nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	if result := h.DB.Where("IATA_code = ?", body.Origin).First(&origin); result.Error != nil {
		jsonData := []byte(`{"error": "Origin airport not found"}`)
		c.AbortWithStatusJSON(http.StatusNotFound, jsonData)
		return
	}

	flight.Origin = origin

	if result := h.DB.Where("IATA_code = ?", body.Destination).First(&destination); result.Error != nil {
		jsonData := []byte(`{"error": "Destination airport not found"}`)
		c.AbortWithStatusJSON(http.StatusNotFound, jsonData)
		return
	}

	flight.Destination = destination

	number, err := strconv.Atoi(body.Number)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	flight.Number = number

	if result := h.DB.First(&flight, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	flight.Number = number
	flight.FlightClasses = body.FlightClasses

	h.DB.Save(&flight)

	c.JSON(http.StatusOK, &flight)
}
