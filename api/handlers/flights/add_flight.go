package flights

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type AddFlightBodyRequest struct {
	OriginID      uint                 `json:"origin_id" binding:"required"`
	DestinationID uint                 `json:"destination_id" binding:"required"`
	FlightClasses []models.FlightClass `json:"flight_classes" binding:"required"`
	Number        string               `json:"number" binding:"required"`
}

func (h handler) AddFlight(c *gin.Context) {
	body := AddFlightBodyRequest{}

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

	if result := h.DB.First(&origin, body.OriginID); result.Error != nil {
		jsonData := []byte(`{"error": "Origin airport not found"}`)
		c.AbortWithStatusJSON(http.StatusNotFound, jsonData)
		return
	}

	flight.OriginID = origin.ID

	if result := h.DB.First(&destination, body.DestinationID); result.Error != nil {
		jsonData := []byte(`{"error": "Destination airport not found"}`)
		c.AbortWithStatusJSON(http.StatusNotFound, jsonData)
		return
	}

	flight.DestinationID = destination.ID

	number, err := strconv.Atoi(body.Number)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	flight.Number = number
	flight.FlightClasses = body.FlightClasses

	if result := h.DB.Create(&flight); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &flight)
}
