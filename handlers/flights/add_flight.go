package flights

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/models"
)

type AddFlightBodyRequest struct {
	OriginID      uint                 `json:"origin_id" binding:"required" example:"1" format:"integer" description:"ID of the origin airport"`
	DestinationID uint                 `json:"destination_id" binding:"required" example:"2" format:"integer" description:"ID of the destination airport"`
	FlightClasses []models.FlightClass `json:"flight_classes" binding:"required" description:"Array of Flight Classes"`
	Number        string               `json:"number" binding:"required" example:"FA123" format:"string" description:"The flight number"`
}

// @Summary Add a new flight
// @Description Add a new flight into the database.
// @Tags Flights
// @Accept  json
// @Produce  json
// @Param origin_id body uint true "ID of the origin airport"
// @Param destination_id body uint true "ID of the destination airport"
// @Param flight_classes body array true "Array of flight classes. Each element is a FlightClass object."
// @Param number body string true "Flight number"
// @Success 201 {object} models.Flight "Successfully added the flight"
// @Failure 400 {object} error "Invalid request body"
// @Failure 404 {object} error "Origin/Destination airport not found"
// @Failure 409 {object} error "Flight number conflict"
// @Router /flights [post]
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
